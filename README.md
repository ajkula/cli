# cli
a toolkit CLI app in Golang. lot of work to be done...

<pre>Usage: cli [options]
Options:
  -a, --ascii string
    	Display ascii art from local images
  -c, --category string
    	Search News by category
        Usage: cli -n [ISO 3166-1 alpha-2 country code] -c {one of:}
        [business entertainment general health science sports technology]
  -C, --com string
    	Search Reddit comments by postId
        Usage: cli -R [reddit keyword] -C [postId]

  -e, --env string
    	Display the env as key/val
  -i, --ip string
    	Remote Network details
  -m, --movie string
    	Search Movies
  -N, --net string
    	List local Network available adresses
  -n, --news string
    	Search News by country code (ex: fr, us)
  -p, --project string
    	Create a Node.js micro-service by a name
        Usage: cli -p [project name]
        to use in terminal emulator under win env

  -P, --publi string
    	Find scientific publications by search-word
        Usage: cli -P [search term]

  -R, --reddit string
    	Search Reddit posts by keyword
  -r, --repo string
    	Search Github repos by User
        Usage: cli -u [user name] -r &apos;y&apos;

  -u, --user string
    	Search Github Users
  -w, --weather string
    	get weather by [city,country code] (ex: paris,fr)
  -x, --x string
    	Width in chars of displayed ascii images</pre>
