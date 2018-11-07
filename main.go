// "Package main" is the namespace declaration
// "main" is a keyword that tells GO that this project is intended to run as a binary/executable (as opposed to a Library)
package main

// importing standard libraries & third party library
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"time"

	// aliasing library names
	flag "github.com/ogier/pflag"
)

func check(e error) {
	if e != nil {
		log.Fatalf("Error retrieving data: %s\n", e)
	}
}

type Folders struct {
	currentFolder string
	controllers   string
	connectors    string
	models        string
}

type Filenames struct {
	serverFile             string
	indexFile              string
	abstractModelFile      string
	testControllerFile     string
	healthControllerFile   string
	abstractControllerFile string
	packageJson            string
	storeMock              string
	readme                 string
	gitignore              string
	empty                  string
}

// flags
var (
	user      string
	repo      string
	movie     string
	genre     string
	news      string
	category  string
	reddit    string
	com       string
	proj      string
	dir       string
	folders   Folders
	filenames Filenames
)

// "main" is the entry point of our CLI app
func main() {
	// parse flags
	flag.Parse()

	// if user does not supply flags, print usage
	if flag.NFlag() == 0 {
		printUsage()
	}

	if proj != "" {
		proj := cleanQuotes(proj)
		folders.currentFolder = "." + dir + "/" + proj + "/"
		folders.connectors = folders.currentFolder + "connectors/"
		folders.controllers = folders.currentFolder + "controllers/"
		folders.models = folders.currentFolder + "models/"

		filenames.gitignore = ".gitignore"
		filenames.abstractModelFile = "AbstractModel.js"
		filenames.abstractControllerFile = "Abstract.js"
		filenames.healthControllerFile = "HealthController.js"
		filenames.indexFile = "index.js"
		filenames.packageJson = "package.json"
		filenames.readme = "README.md"
		filenames.serverFile = "Server.js"
		filenames.storeMock = "store-mock.json"
		filenames.testControllerFile = "testController.js"
		filenames.empty = "EMPTY"

		folders.write(proj)
	}

	if reddit != "" {
		reddit := cleanQuotes(reddit)
		if com != "" {
			comm := cleanQuotes(com)
			coms := getRedditComments(comm)
			fmt.Printf("Searching reddit comments ID: %s\n", comm)
			for _, res := range coms {
				for _, result := range res.Data.Children {
					if result.Data.Selftext != "" {
						{
							fmt.Println(`Date:                `, result.Data.CreatedUTC)
							fmt.Println(`Author:              `, result.Data.Author)
							fmt.Println(`PostId:              `, result.Data.ID)
							fmt.Println(`PostContent:         `, result.Data.Selftext)
							fmt.Println(`*************************** Post ***************************`)
						}
					} else if result.Data.Body != "" {
						{
							fmt.Println(`Date:                `, result.Data.CreatedUTC)
							fmt.Println(`Author:              `, result.Data.Author)
							fmt.Println(`PostId:              `, result.Data.ID)
							fmt.Println(`CommentContent:      `, result.Data.Body)
							fmt.Println(`************************ Comments **************************`)
						}
					}
				}
			}
		} else {
			fmt.Printf("Searching reddit post(s): %s\n", reddit)
			posts := getRedditPosts(reddit)
			for _, result := range posts.Data.Children {
				if result.Data.Selftext != "" {
					fmt.Printf("Searching reddit post(s): %s\n", com)
					{
						fmt.Println(`Date:                `, result.Data.CreatedUTC)
						fmt.Println(`Author:              `, result.Data.Author)
						fmt.Println(`PostId:              `, result.Data.ID)
						fmt.Println(`PostContent:         `, result.Data.Selftext)
						fmt.Println(`************************** Posts ***************************`)
					}
				}
			}
		}
	}

	// if multiple users are passed separated by commas, store them in a "users" array
	if movie != "" {
		movies := strings.Split(cleanQuotes(movie), ",")
		fmt.Printf("Searching movie(s): %s\n", strings.Split(movie, ","))
		if len(movies) > 0 {
			for _, u := range movies {
				result := getMovie(u)
				fmt.Println(`Title:         `, result.Title)
				fmt.Println(`Year:          `, result.Year)
				fmt.Println(`Type:          `, result.Type)
				fmt.Println(`Rated:         `, result.Rated)
				fmt.Println(`Released:      `, result.Released)
				fmt.Println(`Runtime:       `, result.Runtime)
				fmt.Println(`Genre:         `, result.Genre)
				fmt.Println(`Director:      `, result.Director)
				fmt.Println(`Writer:        `, result.Writer)
				fmt.Println(`Actors:        `, result.Actors)
				fmt.Println(`Plot:          `, result.Plot)
				fmt.Println(`Language:      `, result.Language)
				fmt.Println(`Country:       `, result.Country)
				fmt.Println(`Awards:        `, result.Awards)
				fmt.Println(`Poster:        `, result.Poster)
				fmt.Println(`imdbRating:    `, result.ImdbRating)
				fmt.Println(`ImdbVotes:     `, result.ImdbVotes)
				fmt.Println(`DVD:           `, result.DVD)
				fmt.Println(`ID:            `, result.ID)
				fmt.Println("")
			}
		}
	}

	// if multiple users are passed separated by commas, store them in a "users" array
	if user != "" {
		users := strings.Split(user, ",")
		if repo != "" {
			fmt.Printf("Searching [%s]'s repo(s): \n", user)
			res := getRepos(user)
			for _, result := range res.Repos {
				fmt.Println("****************************************************")
				fmt.Println(`Name:              `, result.Name)
				fmt.Println(`Private:           `, result.Private)
				// fmt.Println(`Html_url:          `, result.Html_url)
				fmt.Println(`Description:       `, result.Description)
				// fmt.Println(`Created_at:        `, result.CreatedAt)
				fmt.Println(`Updated_at:        `, result.UpdatedAt)
				fmt.Println(`Git_url:           `, result.GitURL)
				fmt.Println(`Size:              `, result.Size)
				fmt.Println(`Language:          `, result.Language)
				// fmt.Println(`Open_issues_count: `, result.Open_issues_count)
				// fmt.Println(`Forks:             `, result.Forks)
				// fmt.Println(`Watchers:          `, result.Watchers)
				// fmt.Println(`Default_branch:    `, result.Default_branch)
				fmt.Println(`ID:                `, result.Id)
			}
		} else {
			fmt.Printf("Searching user(s): %s\n", users)
			if len(users) > 0 {
				for _, u := range users {
					result := getUsers(u)
					fmt.Println(`Username:        `, result.Login)
					fmt.Println(`Name:            `, result.Name)
					fmt.Println(`Email:           `, result.Email)
					fmt.Println(`Bio:             `, result.Bio)
					fmt.Println(`Location:        `, result.Location)
					fmt.Println(`CreatedAt:       `, result.CreatedAt)
					fmt.Println(`UpdatedAt:       `, result.UpdatedAt)
					fmt.Println(`ReposURL:        `, result.ReposURL)
					fmt.Println(`Followers:       `, result.Followers)
					fmt.Println(`GistsURL:        `, result.GistsURL)
					fmt.Println(`Hireable:        `, result.Hireable)
					fmt.Println("******************* Statistics *********************")
					if len(result.Stats) > 0 {
						for stat, i := range result.Stats {
							x := strings.Repeat(" ", 29-len(stat+strconv.Itoa(i)))
							// y := strings.Repeat(" ", 3-len(strconv.Itoa(i)))
							fmt.Println(`*      ` + stat + x + strconv.Itoa(i) + " %")
						}
					}
					fmt.Println("****************************************************")
				}
			}
		}
	}

	// if multiple users are passed separated by commas, store them in a "users" array
	if news != "" {
		fmt.Printf("Getting news: %s\n", news)
		results := getNews(news)
		// fmt.Println(results)
		for _, res := range results.Articles {
			fmt.Println("**********************************************************")
			fmt.Println(`Source:             `, res.Source.Name)
			fmt.Println(`Publishing date:    `, res.PublishedAt)
			fmt.Println(`Title:              `, res.Title)
			// fmt.Println(`Description:        `, res.Description)
			fmt.Println(`Content:            `, res.Content)
			fmt.Println(`Url:                `, res.Url)
			fmt.Println(`UrlToImage:         `, res.UrlToImage)
		}
	}
}

// "for... range" loop in GO allows us to iterate over each element of the array.
// "range" keyword can return the index of the element (e.g. 0, 1, 2, 3 ...etc)
// and it can return the actual value of the element.
// Since GO does not allow unused variables, we use the "_" character to tell GO we don't care about the index, but
// we want to get the actual user we're looping over to pass to the function.

func GetDateFromTimeStamp(dtformat string) string {
	form := "Jan 2, 2006 at 3:04 PM"
	t2, _ := time.Parse(form, dtformat)
	return t2.Format("20060102150405")
}

// "init" is a special function. GO will execute the init() function before the main.
func init() {
	// We pass the user variable we declared at the package level (above).
	// The "&" character means we are passing the variable "by reference" (as opposed to "by value"),
	// meaning: we don't want to pass a copy of the user variable. We want to pass the original variable.
	flag.StringVarP(&user, "user", "u", "", "Search Github Users")
	flag.StringVarP(&repo, "repo", "r", "", "Search Github repos by User")
	flag.StringVarP(&movie, "movie", "m", "", "Search Movies")
	flag.StringVarP(&genre, "genre", "g", "", "Search Movie by genre")
	flag.StringVarP(&news, "news", "n", "", "Search News by country ode (ex: fr, us)")
	flag.StringVarP(&category, "category", "c", "", "Search News by category")
	flag.StringVarP(&reddit, "reddit", "R", "", "Search Reddit posts by keyword")
	flag.StringVarP(&com, "com", "C", "", "Search Reddit comments by postId")
	flag.StringVarP(&proj, "project", "p", "", "Create a Node.js micro-service by a name")

	dir, _ := syscall.Getwd()
	fmt.Println("dossier courant:", dir)
	// project()
	// fmt.Println(createProject("SANDBOX"))
}

// printUsage is a custom function we created to print usage for our CLI app
func printUsage() {
	fmt.Printf("Usage: %s [options]\n", os.Args[0])
	fmt.Println("Options:")
	flag.PrintDefaults()
	os.Exit(1)
}

func (pr Folders) write(name string) {

	cmd := exec.Command("ls")
	cmd.Start()
	cmd = exec.Command("mkdir", name)
	cmd.Start()
	cmd = exec.Command("cd", name)
	cmd.Start()
	cmd = exec.Command("mkdir", pr.connectors)
	cmd.Start()
	cmd = exec.Command("mkdir", pr.controllers)
	cmd.Start()
	cmd = exec.Command("mkdir", pr.currentFolder)
	cmd.Start()
	cmd = exec.Command("mkdir", pr.models)
	cmd.Start()

	packageJson, err := os.Create(pr.currentFolder + filenames.packageJson)
	fmt.Println("os create:", pr.currentFolder+filenames.packageJson)
	check(err)
	defer packageJson.Close()
	pjs := bufio.NewWriter(packageJson)
	b, err := pjs.WriteString(createProject(name).packageJson)
	check(err)
	fmt.Println("wrote "+pr.currentFolder+filenames.packageJson, string(b)+" bytes")
	pjs.Flush()

	indexFile, err := os.Create(pr.currentFolder + filenames.indexFile)
	check(err)
	defer indexFile.Close()
	idx := bufio.NewWriter(indexFile)
	b, err = idx.WriteString(createProject(name).indexFile)
	check(err)
	fmt.Println("wrote "+pr.currentFolder+filenames.indexFile, string(b)+" bytes")
	idx.Flush()

	gitignore, err := os.Create(pr.currentFolder + filenames.gitignore)
	check(err)
	defer gitignore.Close()
	git := bufio.NewWriter(gitignore)
	b, err = git.WriteString(createProject(name).gitignore)
	check(err)
	fmt.Println("wrote "+pr.currentFolder+filenames.gitignore, string(b)+" bytes")
	git.Flush()

	readme, err := os.Create(pr.currentFolder + filenames.readme)
	check(err)
	defer readme.Close()
	rdm := bufio.NewWriter(readme)
	b, err = rdm.WriteString(createProject(name).readme)
	check(err)
	fmt.Println("wrote "+pr.currentFolder+filenames.readme, string(b)+" bytes")
	rdm.Flush()

	serverFile, err := os.Create(pr.currentFolder + filenames.serverFile)
	check(err)
	defer serverFile.Close()
	srv := bufio.NewWriter(serverFile)
	b, err = srv.WriteString(createProject(name).serverFile)
	check(err)
	fmt.Println("wrote "+pr.currentFolder+filenames.serverFile, string(b)+" bytes")
	srv.Flush()

	storeMock, err := os.Create(pr.currentFolder + filenames.storeMock)
	check(err)
	defer storeMock.Close()
	str := bufio.NewWriter(storeMock)
	b, err = str.WriteString(createProject(name).storeMock)
	check(err)
	fmt.Println("wrote "+pr.currentFolder+filenames.storeMock, string(b)+" bytes")
	str.Flush()

	empty, err := os.Create(pr.connectors + filenames.empty)
	check(err)
	defer empty.Close()
	ept := bufio.NewWriter(empty)
	b, err = ept.WriteString(createProject(name).empty)
	check(err)
	fmt.Println("wrote "+pr.currentFolder+filenames.empty, string(b)+" bytes")
	ept.Flush()

	abstractControllerFile, err := os.Create(pr.controllers + filenames.abstractControllerFile)
	check(err)
	defer abstractControllerFile.Close()
	abc := bufio.NewWriter(abstractControllerFile)
	b, err = abc.WriteString(createProject(name).abstractControllerFile)
	check(err)
	fmt.Println("wrote "+pr.currentFolder+filenames.abstractControllerFile, string(b)+" bytes")
	abc.Flush()

	healthControllerFile, err := os.Create(pr.controllers + filenames.healthControllerFile)
	check(err)
	defer healthControllerFile.Close()
	hlc := bufio.NewWriter(healthControllerFile)
	b, err = hlc.WriteString(createProject(name).healthControllerFile)
	check(err)
	fmt.Println("wrote "+pr.currentFolder+filenames.healthControllerFile, string(b)+" bytes")
	hlc.Flush()

	testControllerFile, err := os.Create(pr.controllers + filenames.testControllerFile)
	check(err)
	defer testControllerFile.Close()
	tst := bufio.NewWriter(testControllerFile)
	b, err = tst.WriteString(createProject(name).testControllerFile)
	check(err)
	fmt.Println("wrote "+pr.currentFolder+filenames.testControllerFile, string(b)+" bytes")
	tst.Flush()

	cmd = exec.Command("npm", "i")
	cmd.Dir = name
	out, err := cmd.Output()
	check(err)
	fmt.Println(string(out))

	cmd = exec.Command("explorer", ".")
	cmd.Dir = name
	cmd.Start()
}
