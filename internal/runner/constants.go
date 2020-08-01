package runner

const (
	email       = "secinfo@kitabisa.com"
	version     = "0.0.1-dev3"
	development = true
	banner      = `
	  __      __       
	 / /____ / /__ ____
	/ __/ -_) / -_) __/
	\__/\__/_/\__/_/   
	                v` + version

	usage = `  [buffers] | teler [options]
  teler -c [...] [options]`
	example = `  teler -c [...] -i /var/log/nginx/access.log
  [kubectl logs|tail -f|cat] ... | teler -c [...] -x 25
  `
)
