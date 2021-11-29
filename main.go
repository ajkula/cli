// "Package main" is the namespace declaration
// "main" is a keyword that tells GO that this project is intended to run as a binary/executable (as opposed to a Library)
package main

// importing standard libraries & third party library
import (
	"bufio"
	"fmt"
	"math"
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
		fmt.Printf("Error retrieving data: %s\n", e)
	}
}

// Folders struct
type Folders struct {
	currentFolder string
	controllers   string
	connectors    string
	models        string
	test          string
	public        string
}

// Filenames struct
type Filenames struct {
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

// flags
var (
	x         string
	ip        string
	img       string
	netw      string
	city      string
	user      string
	publi     string
	repo      string
	movie     string
	genre     string
	news      string
	category  string
	reddit    string
	com       string
	proj      string
	dir       string
	osTool    string
	docker    string
	folders   Folders
	filenames Filenames
)

// "main" is the entry point of our CLI app
func main() {
	var argsLength int
	if len(os.Args) > 0 {
		argsLength = len(os.Args)
		argsContent := ""
		for i := 0; i < argsLength; i++ {
			argsContent += os.Args[i] + " "
		}
		fmt.Println(argsContent)
	}

	// parse flags
	flag.Parse()

	// if user does not supply flags, print usage
	if flag.NFlag() == 0 {
		printUsage()
	}
	if publi != "" {
		DisplayPublications(publi)
	}

	if osTool != "" {
		ListOSTools()
	}

	if docker == "l" || docker == "list" {
		ListContainer()
	}
	// ReadSettingsFile()
	listLocalAddresses(netw, ip)

	if proj != "" {
		proj := cleanQuotes(proj)
		folders.currentFolder = "." + dir + "/" + proj + "/"
		folders.connectors = folders.currentFolder + "connectors/"
		folders.controllers = folders.currentFolder + "controllers/"
		folders.models = folders.currentFolder + "models/"
		folders.test = folders.currentFolder + "test/"
		folders.public = folders.currentFolder + "public/"

		filenames.gitignore = ".gitignore"
		filenames.abstractModelFile = "AbstractModel.js"
		filenames.abstractControllerFile = "Abstract.js"
		filenames.healthControllerFile = "HealthController.js"
		filenames.indexFile = "index.js"
		filenames.packageJSON = "package.json"
		filenames.readme = "README.md"
		filenames.serverFile = "Server.js"
		filenames.storeMock = "store-mock.json"
		filenames.testControllerFile = "testController.js"
		filenames.apiTests = "apiTests.js"
		filenames.empty = "EMPTY"
		folders.write(proj)
	}

	if reddit != "" {
		reddit := cleanQuotes(reddit)
		if com != "" {
			com := cleanQuotes(com)
			coms := getRedditComments(com)
			fmt.Printf("Searching reddit comments ID: %s\n", com)
			for _, res := range coms {
				for _, result := range res.Data.Children {
					if result.Data.Selftext != "" {
						{
							fmt.Println(`Date:                `, GetDateFromTimeStamp(result.Data.CreatedUTC))
							fmt.Println(`Author:              `, result.Data.Author)
							fmt.Println(`PostId:              `, result.Data.ID)
							fmt.Println(`PostContent:         `, result.Data.Selftext)
							fmt.Println(`*************************** Post ***************************`)
						}
					} else if result.Data.Body != "" {
						{
							fmt.Println(`Date:                `, GetDateFromTimeStamp(result.Data.CreatedUTC))
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
						fmt.Println(`Date:                `, GetDateFromTimeStamp(result.Data.CreatedUTC))
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
		DisplayMoviesByName(movie)
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
				// fmt.Println(`HTMLURL:          `, result.HTMLURL)
				fmt.Println(`Description:       `, result.Description)
				// fmt.Println(`Created_at:        `, result.CreatedAt)
				fmt.Println(`Updated_at:        `, result.UpdatedAt)
				fmt.Println(`Git_url:           `, result.GitURL)
				fmt.Println(`Size:              `, result.Size)
				fmt.Println(`Language:          `, result.Language)
				// fmt.Println(`Open_issues_count: `, result.Open_issues_count)
				// fmt.Println(`Forks:             `, result.Forks)
				// fmt.Println(`Watchers:          `, result.Watchers)
				// fmt.Println(`DefaultBranch:    `, result.DefaultBranch)
				fmt.Println(`ID:                `, result.ID)
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
							formatSpacedStringWithItoa(stat, i)
						}
					}
					fmt.Println("****************************************************")
				}
			}
		}
	}

	// if multiple users are passed separated by commas, store them in a "users" array
	if news != "" {
		DisplayNews(news, category, x)
	}

	if img != "" {
		DisplayASCIIFromLocalFile(img, x)
	}

	if city != "" {
		city = cleanQuotes(city)
		DisplayWeather(city)
	}
}

func formatSpacedStringWithItoa(str string, i int) {
	formatSpacedStrings(str, strconv.Itoa(i), `*      `, " %")
}

func Abs(x int) int {
	if x == math.MinInt64 {
		return math.MaxInt64
	}
	if x < 0 {
		return -x
	}
	return x
}

func formatSpacedStrings(strA, strB string, arr ...string) {
	left := ""
	right := ""
	defaultLength := 29

	if len(arr) > 0 {
		left = arr[0]
	}

	if len(arr) > 1 {
		right = arr[1]
	}

	if len(arr) > 2 {
		val, err := strconv.Atoi(arr[2])
		if err == nil {
			defaultLength = val
		}
	}

	x := strings.Repeat(" ", Abs(defaultLength-len(strA)))
	fmt.Println(left + strA + x + strB + right)

	//  x := strings.Repeat(" ", 29-len(stat+strconv.Itoa(i)))
	//  fmt.Println(`*      ` + stat + x + strconv.Itoa(i) + " %")
}

// "for... range" loop in GO allows us to iterate over each element of the array.
// "range" keyword can return the index of the element (e.g. 0, 1, 2, 3 ...etc)
// and it can return the actual value of the element.
// Since GO does not allow unused variables, we use the "_" character to tell GO we don't care about the index, but
// we want to get the actual user we're looping over to pass to the function.

// GetDateFromTimeStamp function returns time struct from float64 data
func GetDateFromTimeStamp(dtformat float64) time.Time {
	return time.Unix(int64(dtformat), 0)
}

func cleanTags(text string) string {
	br := "<br />"
	return strings.Replace(text, br, "\n", -1)
}

// "init" is a special function. GO will execute the init() function before the main.
func init() {
	// We pass the user variable we declared at the package level (above).
	// The "&" character means we are passing the variable "by reference" (as opposed to "by value"),
	// meaning: we don't want to pass a copy of the user variable. We want to pass the original variable.
	flag.StringVarP(&city, "weather", "w", "", "get weather by [city,country code] (ex: paris,fr)")
	flag.StringVarP(&user, "user", "u", "", "Search Github Users")
	flag.StringVarP(&repo, "repo", "r", "", "Search Github repos by User\n        Usage: cli -u [user name] -r 'y'\n")
	flag.StringVarP(&movie, "movie", "m", "", "Search Movies")
	// flag.StringVarP(&genre, "genre", "g", "", "Search Movie by genre\n        Usage: cli -g {not yet implemented}\n")
	flag.StringVarP(&news, "news", "n", "", "Search News by country code (ex: fr, us)")
	flag.StringVarP(&category, "category", "c", "", "Search News by category\n        Usage: cli -n [ISO 3166-1 alpha-2 country code] -c {one of:}\n        [business entertainment general health science sports technology]")
	flag.StringVarP(&reddit, "reddit", "R", "", "Search Reddit posts by keyword")
	flag.StringVarP(&com, "com", "C", "", "Search Reddit comments by postId\n        Usage: cli -R [reddit keyword] -C [postId]\n")
	flag.StringVarP(&proj, "project", "p", "", "Create a Node.js micro-service by a name\n        Usage: cli -p [project name]\n        to use in terminal emulator under win env\n")
	flag.StringVarP(&publi, "publi", "P", "", "Find scientific publications by search-word\n        Usage: cli -P [search term]\n")
	flag.StringVarP(&osTool, "env", "e", "", "Display the env as key/val")
	flag.StringVarP(&docker, "docker", "d", "", "Docker tool\n        Usage: cli -d [list/l]\n")
	flag.StringVarP(&x, "x", "x", "", "Width in chars of displayed ascii images")
	flag.StringVarP(&netw, "net", "N", "", "List local Network available adresses")
	flag.StringVarP(&ip, "ip", "i", "", "Remote Network details")
	flag.StringVarP(&img, "ascii", "a", "", "Display ascii art from local images")

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
	cmd = exec.Command("mkdir", pr.public)
	cmd.Start()
	cmd = exec.Command("mkdir", pr.models)
	cmd.Start()
	cmd = exec.Command("mkdir", pr.test)
	cmd.Start()
	cmd = exec.Command("mkdir", pr.connectors)
	cmd.Start()
	cmd = exec.Command("mkdir", pr.controllers)
	cmd.Start()
	cmd = exec.Command("mkdir", pr.currentFolder)
	cmd.Start()

	packageJSON, err := os.Create(pr.currentFolder + filenames.packageJSON)
	fmt.Println("os create:", pr.currentFolder+filenames.packageJSON)
	check(err)
	defer packageJSON.Close()
	pjs := bufio.NewWriter(packageJSON)
	b, err := pjs.WriteString(createProject(name).packageJSON)
	check(err)
	fmt.Println("wrote "+pr.currentFolder+filenames.packageJSON, intToString(b)+" bytes")
	pjs.Flush()

	indexFile, err := os.Create(pr.currentFolder + filenames.indexFile)
	check(err)
	defer indexFile.Close()
	idx := bufio.NewWriter(indexFile)
	b, err = idx.WriteString(createProject(name).indexFile)
	check(err)
	fmt.Println("wrote "+pr.currentFolder+filenames.indexFile, intToString(b)+" bytes")
	idx.Flush()

	gitignore, err := os.Create(pr.currentFolder + filenames.gitignore)
	check(err)
	defer gitignore.Close()
	git := bufio.NewWriter(gitignore)
	b, err = git.WriteString(createProject(name).gitignore)
	check(err)
	fmt.Println("wrote "+pr.currentFolder+filenames.gitignore, intToString(b)+" bytes")
	git.Flush()

	readme, err := os.Create(pr.currentFolder + filenames.readme)
	check(err)
	defer readme.Close()
	rdm := bufio.NewWriter(readme)
	b, err = rdm.WriteString(createProject(name).readme)
	check(err)
	fmt.Println("wrote "+pr.currentFolder+filenames.readme, intToString(b)+" bytes")
	rdm.Flush()

	serverFile, err := os.Create(pr.currentFolder + filenames.serverFile)
	check(err)
	defer serverFile.Close()
	srv := bufio.NewWriter(serverFile)
	b, err = srv.WriteString(createProject(name).serverFile)
	check(err)
	fmt.Println("wrote "+pr.currentFolder+filenames.serverFile, intToString(b)+" bytes")
	srv.Flush()

	storeMock, err := os.Create(pr.currentFolder + filenames.storeMock)
	check(err)
	defer storeMock.Close()
	str := bufio.NewWriter(storeMock)
	b, err = str.WriteString(createProject(name).storeMock)
	check(err)
	fmt.Println("wrote "+pr.currentFolder+filenames.storeMock, intToString(b)+" bytes")
	str.Flush()

	apiTests, err := os.Create(pr.test + filenames.apiTests)
	fmt.Println("os create:", pr.test+filenames.apiTests)
	check(err)
	defer apiTests.Close()
	apt := bufio.NewWriter(apiTests)
	b, err = apt.WriteString(createProject(name).apiTests)
	check(err)
	fmt.Println("wrote "+pr.test+filenames.apiTests, intToString(b)+" bytes")
	apt.Flush()

	empty, err := os.Create(pr.connectors + filenames.empty)
	check(err)
	defer empty.Close()
	ept := bufio.NewWriter(empty)
	b, err = ept.WriteString(createProject(name).empty)
	check(err)
	fmt.Println("wrote "+pr.connectors+filenames.empty, intToString(b)+" bytes")
	ept.Flush()

	abstractControllerFile, err := os.Create(pr.controllers + filenames.abstractControllerFile)
	check(err)
	defer abstractControllerFile.Close()
	abc := bufio.NewWriter(abstractControllerFile)
	b, err = abc.WriteString(createProject(name).abstractControllerFile)
	check(err)
	fmt.Println("wrote "+pr.controllers+filenames.abstractControllerFile, intToString(b)+" bytes")
	abc.Flush()

	healthControllerFile, err := os.Create(pr.controllers + filenames.healthControllerFile)
	check(err)
	defer healthControllerFile.Close()
	hlc := bufio.NewWriter(healthControllerFile)
	b, err = hlc.WriteString(createProject(name).healthControllerFile)
	check(err)
	fmt.Println("wrote "+pr.controllers+filenames.healthControllerFile, intToString(b)+" bytes")
	hlc.Flush()

	testControllerFile, err := os.Create(pr.controllers + filenames.testControllerFile)
	check(err)
	defer testControllerFile.Close()
	tst := bufio.NewWriter(testControllerFile)
	b, err = tst.WriteString(createProject(name).testControllerFile)
	check(err)
	fmt.Println("wrote "+pr.controllers+filenames.testControllerFile, intToString(b)+" bytes")
	tst.Flush()

	abstractModelFile, err := os.Create(pr.models + filenames.abstractModelFile)
	check(err)
	defer abstractModelFile.Close()
	abm := bufio.NewWriter(abstractModelFile)
	b, err = abm.WriteString(createProject(name).abstractModelFile)
	check(err)
	fmt.Println("wrote "+pr.models+filenames.abstractModelFile, intToString(b)+" bytes")
	abm.Flush()

	cmd = exec.Command("npm", "i")
	cmd.Dir = name
	out, err := cmd.Output()
	check(err)
	fmt.Println(string(out))
	cmd = exec.Command("npm", "run", "test")
	cmd.Dir = name
	out, err = cmd.Output()
	check(err)
	fmt.Println(string(out))

	cmd = exec.Command("explorer", ".")
	cmd.Dir = name
	out, err = cmd.Output()
	check(err)
	fmt.Println(string(out))
	cmd.Start()
}

func intToString(num int) string {
	return fmt.Sprintf("%d", num)
}

// func writeFiles(pr Folders, fileName string) {
// 	file, err := os.Create(pr.currentFolder + filename)
// 	check(err)
// 	defer file.Close()
// 	writer := bufio.NewWriter(file)
// 	b, err = writer.WriteString(createProject(name).(fileName))
// 	check(err)
// 	fmt.Println("wrote "+pr.currentFolder+filename, string(b)+" bytes")
// 	writer.Flush()
// 	wg.Done()
// }
