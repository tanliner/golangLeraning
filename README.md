### Go Lang Learning

Is there is a lang like C/C++ which has high performance ?
Is there is a lang like Java which has a Garbage Collection period ?
Is there is a lang like Python which is easy to learning, but it has a static type check,
function programming, concurrent programming?

Now, Go is ready to solve problem above.

### CSP (Communication Sequential Process)
you don't worry about the LOCK/CallBack, It was solve by sub system 
the `NODE` will calculate concurrently

### compile
go build
go run

it will compile into machine code, faster

### fmt library
this format library

like C, `printf("this is a int %d", 100)`, in Go, `fmt.Printf(this is a int %d", 100)`, see

### net library
`import (
 	"net/http"
 )`
 

 ### dep
the package dependency, like gradle lib
use this command `brew install dep` , and it will automatically to install the go latest version.

### Nginx
download Nginx
[Install-with-brew](https://coderwall.com/p/dgwwuq/installing-nginx-in-mac-os-x-maverick-with-homebrew)
the config is here: `/usr/local/etc/nginx/nginx.conf`

to review [this blog](https://www.jianshu.com/p/33d4a3fdc483) to check Nginx config

append the extra virtual server below default server
```
    server {
        listen       8080;
        server_name  localhost;
        location / {
            root   html;
            index  index.html index.htm;
        }
        error_page   500 502 503 504  /50x.html;
        ...
    }
    // the request to 8091 will redirect to localhost:9090;
    server {
        listen 8091;
        root   html;
        index  index.html index.htm index.php;
        # send request back to your server
        location / {
            # go, the local host
            proxy_pass  http://localhost:9090;
       }
    }
```

Some nginx command maybe useful
`nginx`
`nginx -s stop`
`nginx -s reload`
`nginx -h`

 
### go https
##### generate certificate file
Key considerations for algorithm "RSA" ≥ 2048-bit<br/>
```openssl genrsa -out server.key 2048```

Key considerations for algorithm "ECDSA" ≥ secp384r1<br/>
List ECDSA the supported curves (openssl ecparam -list_curves)<br/>
```openssl ecparam -genkey -name secp384r1 -out server.key```

Generation of self-signed(x509) public key (PEM-encodings .pem|.crt) based on the private (.key)<br/>
```openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650```

##### simple GO https/TLS server
file: `src/nginx/simple/https`

>TLS Server

[Ref-Github-Gist](https://gist.github.com/denji/12b3a568f092ab951456)
`src/nginx/server`<br/>
`src/nginx/client`<br/>

### Note
the Go file name rule I do not know yet

the file tree of Go not like Java
src & bin & pkg

the `src` is the All project source codes