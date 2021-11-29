package main

// Project struct data
type Project struct {
	serverFile             string
	indexFile              string
	abstractModelFile      string
	testControllerFile     string
	healthControllerFile   string
	abstractControllerFile string
	apiTests               string
	packageJSON            string
	storeMock              string
	readme                 string
	gitignore              string
	empty                  string
}

/* func project() {
	cmd := exec.Command("go", "version")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)
} */

func createProject(name string) Project {
	var project Project
	project.empty = ``
	project.serverFile = `
const express = require('express');
const test = require('./controllers/testController');
const fs = require('fs');

module.exports = class Server {

    /**
     * Init project
     */
    constructor () {
      this.container = {user: "greg"};
      this.app = express();
      this.app.use(express.json());
	  this.app.use('/static', express.static(__dirname + '/public'));
      this.loadDBs();
      this.loadControllers();
      
    }

  /**
   * Connect to databases
   */
  async loadDBs () {
    // const dynamoConnector = new DynamoConnector('dbName', false)
    // this.container.campaignModel = new CampaignModel(dynamoConnector)
    this.container.storeModel = new Map();
    Object.entries(require('./store-mock.json')).forEach(([k, v]) => {
			this.container.storeModel.set(k, v);
    });
  }

    /**
   * Load Http controllers
   */
  loadControllers () {
    console.log(__dirname + '/controllers/');
    const files = fs.readdirSync(__dirname + '/controllers/');
    files.forEach((file) => {
      if (/Controller\.js/.test(file)) {
        const classCtrl = require(__dirname + '/controllers/' + file);
        this.app.use(new classCtrl(this.container).router);
      }
    })
  }
}
	`
	project.indexFile = `
const env = typeof process.env.NODE_ENV === 'undefined' ? 'development' : process.env.NODE_ENV;
global.__env = env;
global.__DEV = env === 'development';
global.__TEST = env === 'test';
global.__PROD = env === 'production';
global.filePath = __dirname + '/store-mock.json';

const config = {LISTEN_PORT: 3500};
const logule = require('logule').init(module);
const Server = require('./Server');
logule.info("\x1b[31mCONFIG", JSON.stringify(config, null, 6) + "\x1b[0m");

let server = new Server();
let app = server.app;
app.set("PORT", 3500);

let port = Number(parseInt(config.LISTEN_PORT) || 3500);
app.listen(port, function () {
  logule.info('\x1b[36m[EXPRESS] Listening on port ' + app.get("PORT") + "\x1b[0m");
  return true;
});
	`
	project.abstractModelFile = `
	let uuid = require('node-uuid');
	const Promise = require('bluebird');
	const logule = require('logule');
	
	module.exports = class AbstractModel {
		/**
		 * Abstract Database model
		 * @param connector
		 */
		constructor (connector) {
			this.connector = connector;
			this.isTest = process.env.NODE_ENV === 'test';
		}
	
		/**
		 * Fetch a document from Database with ID of model
		 * @param id
		 * @return {Promise.<Map>}
		 */
		findById (id) {
			return this.connector.get(id);
		}
	
		/**
		 * Get collection of documents
		 * @param ids
		 * @return Collection of document
		 */
		async findByMultipleId (ids) {
			let results = {};
			await Promise.map(ids, async (id) => {
				results[id] = await this.connector.get(id);
			});
	
			return results;
		}
	
		/**
		 * Create or update a object in Database
		 * @param id
		 * @param data
		 * @param options
		 * @return {Promise}
		 */
		save (id, data, options = {expiry: 0}) {
			return this.connector.save(id, data, options);
		}
	
		/**
		 * Health check
		 * @return {Promise}
		 */
		async health () {
			const uid = uuid.v4();
			try {
				await this.save(uid, {id: uid, value: '42'}, {'expiry': 20});
				const result = await this.findById(uid);
				return ((result || {}).value || null) === '42';
			} catch (err) {
				if (!this.isTest) logule.error(` + "`Health error: ${err.message}`" + `);
				return false;
			}
		}
	}
	`
	project.packageJSON = `
	{
		"name": "` + name + `",
		"version": "1.0.0",
		"description": "",
		"main": "index.js",
		"scripts": {
			"test": "mocha"
		},
		"author": "",
		"license": "ISC",
		"dependencies": {
			"async": "^2.6.0",
			"bluebird": "^3.5.1",
			"express": "^4.16.2",
			"logule": "^2.1.0",
			"mongodb": "^3.0.3",
			"node-uuid": "^1.4.8"
		},
		"devDependencies": {
			"chai": "^4.3.4",
			"chai-http": "^4.3.0",
			"mocha": "^9.1.3"
		}
	}
	`
	project.abstractControllerFile = `
	let express = require('express')

	module.exports = class AbstractController {
		constructor (container) {
			this.container = container;
			this.router = express.Router();
			this.isTest = process.env.NODE_ENV === 'test';
		}
	
		get (serviceId) {
			if (!this.isTest) console.log("get (" + serviceId + ")")
			return this.container[serviceId];
		}
	}
	`
	project.healthControllerFile = `
	let AbstractController = require('./Abstract')

	module.exports = class HealthController extends AbstractController {
	
		constructor(container) {
			super(container)
	
			this.anyModel = this.get('user');
			// Load routes
			this.router.get('/health', this.checkHealth.bind(this));
		}
	
		/**
		 * Check health
		 * @param req
		 * @param res
		 */
		async checkHealth(req, res) {
			// const status = await this.anyModel // some async call
			res.json({status:'UP'});
		}
	}
	`
	project.testControllerFile = `
	const Abstract = require('./Abstract');
	const logule = require('logule').init(module);
	
	module.exports = class test extends Abstract {
		constructor(container) {
			super(container);
			this.container = container;
			this.router.get('/test/:id', this.getById.bind(this));
			this.router.get('/test', this.message.bind(this));
			this.router.post('/test', this.save.bind(this));
		}

		getById(req, res) {
			const id = req.params.id;
			if (!this.isTest) logule.warn(JSON.stringify([...this.container.storeModel]));
			try {
				const response = Object.assign(this.container.storeModel.get(id), {message: 'ok'});
				if (!this.isTest) logule.warn(response);
				res.json(response);

			} catch (e) {
				if (!this.isTest) logule.error(e.message);
				if (!this.isTest) logule.error(` + "`ID: ${id} is undefined`" + `);
				res.json({status: 404, message: ` + "`${id} undefined`" + `});
			}
		}

		message(req, res) {
			res.json({"Usage:": "/test/:id"});
		}

		save(req, res) {
			const { body, headers } = req;
			if (!body || !body.id)  {
				if (!this.isTest) logule.error(` + "`ID: ${id} is undefined`" + `);
				res.json({status: 404, message: ` + "`${id} undefined`" + `});
			}
			const id = body.id.toString();
			body.id = body.id.toString();
			this.container.storeModel.set(id, body);
			if (!this.isTest) logule.info(` + "`ID: ${id} Recorded!`" + `);
			res.json({ object: this.container.storeModel.get(id), id: id });
		}
	}
	`
	project.storeMock = `
	{
    "1": {
        "id": "1",
        "name": "1 - Kia Motors",
        "click_command": "http://www.google.com",
        "start_date": "2015-10-23",
        "end_date": "2015-11-09",
        "in_pause": true,
        "created_at": "2015-10-25T11:14:42.881Z",
        "updated_at": "2016-09-05T10:08:28.062Z",
        "country": "NOT",
        "message": "ok"
    }
}
	`
	project.readme = `
# sandbox
Bootstrapping NodeJs ES6 Classes server + connectors...
	`
	project.gitignore = `
	# Logs
	logs
	*.log
	npm-debug.log*
	yarn-debug.log*
	yarn-error.log*
	
	# Runtime data
	pids
	*.pid
	*.seed
	*.pid.lock
	
	# Directory for instrumented libs generated by jscoverage/JSCover
	lib-cov
	
	# Coverage directory used by tools like istanbul
	coverage
	
	# nyc test coverage
	.nyc_output
	
	# Grunt intermediate storage (http://gruntjs.com/creating-plugins#storing-task-files)
	.grunt
	
	# Bower dependency directory (https://bower.io/)
	bower_components
	
	# node-waf configuration
	.lock-wscript
	
	# Compiled binary addons (http://nodejs.org/api/addons.html)
	build/Release
	
	# Dependency directories
	node_modules/
	jspm_packages/
	
	# Typescript v1 declaration files
	typings/
	
	# Optional npm cache directory
	.npm
	
	# Optional eslint cache
	.eslintcache
	
	# Optional REPL history
	.node_repl_history
	
	# Output of 'npm pack'
	*.tgz
	
	# Yarn Integrity file
	.yarn-integrity
	
	# dotenv environment variables file
	.env	
	`
	project.apiTests = `
	process.env.NODE_ENV = 'test';

const http = require('http');
const chai = require('chai');
const mocha = require('mocha');
const chaiHttp = require('chai-http');
const logule = require('logule').init(module);
const should = chai.should();
const Server = require('../Server');
chai.use(chaiHttp);
const config = {LISTEN_PORT: 3500};
let server = new Server();
let app = server.app;
const NAME = "John Doe";
const test1User = {
    id: '1',
    name: '1 - Kia Motors',
    click_command: 'http://www.google.com',
    start_date: '2015-10-23',
    end_date: '2015-11-09',
    in_pause: true,
    created_at: '2015-10-25T11:14:42.881Z',
    updated_at: '2016-09-05T10:08:28.062Z',
    country: 'NOT',
    message: 'ok'
  };
const newUser = {
    id: '2',
    name: "John Doe"
}
const test404Get = id => { return {status: 404, message: ` + "`${id} undefined`" + `}};

describe('Api Tests', () => {
    
before(function (done) {
    app.set("PORT", 3500);
            done();
    })
    
    describe('/GET health', () => {
        it('it should GET server health', (done) => {
            chai.request(app.listen(app.get("port")))
                .get('/health')
                .end((err, res) => {
                    res.should.have.status(200);
                    res.body.should.be.a('object');
                    done();
                });
        });
    });

    describe('/GET test', () => {
        it('it should GET existing user 1', (done) => {
            chai.request(app.listen(app.get("port")))
                .get('/test/1')
                .end((err, res) => {
                    res.should.have.status(200);
                    res.body.should.be.a('object');
                    res.body.should.contain(test1User);
                    done();
                });
        });
    });

    describe('/GET test', () => {
        it('it should GET 404 user id 2 undefined', (done) => {
            let id = 3;
            chai.request(app.listen(app.get("port")))
                .get('/test/' + id)
                .end((err, res) => {
                    res.should.have.status(200);
                    res.body.should.be.a('object');
                    res.body.should.contain(test404Get(id));
                    done();
                });
        });
    });

    describe('/POST test', () => {
        it('it should POST new user id: '+newUser.id, (done) => {
            let id = 3;
            chai.request(app.listen(app.get("port")))
                .post('/test')
                .send(newUser)
                .end((err, res) => {
                    res.should.have.status(200);
                    res.body.should.be.a('object');
                    done();
                });
        });
    });

    describe('/GET test', () => {
        it('it should GET new user '+newUser.id, (done) => {
            chai.request(app.listen(app.get("port")))
                .get('/test/'+newUser.id)
                .end((err, res) => {
                    res.should.have.status(200);
                    res.body.should.be.a('object');
                    res.body.should.contain(newUser);
                    done();
                });
        });
    });

});
	`
	return project
}
