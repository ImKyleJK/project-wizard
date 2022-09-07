name | description | languages
--- | --- | ---
NECM | NodeJS, ExpressJS, Connect, MongoDB | javascript | ImKyleJk

# index.js
```
// ====================================================== //
// ================= Import all packages ================ //
// ====================================================== //
const config = require("./config/config.js");
const vhost = require("vhost");
const connect = require("connect");
const express = require("express");
const helmet = require('helmet');
const minifyHTML = require("express-minify-html");
const serveStatic = require("serve-static");
const compression = require("compression");
const cookieParser = require("cookie-parser");
const session = require("express-session");

// ====================================================== //
// ===== Create express app and connect middleware ====== //
// =======================================================//
const app = express();
const con = connect();

// ====================================================== //
// =================== Express configs ================== //
// ====================================================== //
app.use(express.json());
app.use(express.urlencoded({ extended: false }));
app.use(cookieParser('', `domain=${config.domain.domain}`));
app.use(compression());
app.set("view engine", "ejs");
app.use(
    minifyHTML({
        override: true,
        exception_url: false,
        htmlMinifier: {
            removeComments: false, // Dont remove the comments from our EJS pages
            collapseWhitespace: true, // Removes all spaces
            collapseBooleanAttributes: true, // Reduces true & false statements
            removeAttributeQuotes: false, // Remove attribute quotes
            removeEmptyAttributes: true, // Remove empty attributes
            minifyJS: true, // Reduce/Minify JS
        },
    })
);
app.use(
    session({
        secret: config.system.session.secret,
        name: config.system.session.name,
        resave: false,
        saveUninitialized: true,
        cookie: { secure: false },
    })
);
app.use(helmet());
app.use(helmet.hidePoweredBy());
app.use(
    helmet.permittedCrossDomainPolicies({
        permittedPolicies: "all",
    })
);
app.use(helmet.xssFilter());
app.use(helmet.crossOriginResourcePolicy({ policy: "cross-origin" }));
app.use(
    helmet.contentSecurityPolicy({
        useDefaults: true,
        directives: {
            "script-src": ["'self'", `${config.domain.subdomain.cdn}.${config.domain.domain}`, "'unsafe-inline'", "'unsafe-eval'"],
            "style-src": ["'self'", `${config.domain.subdomain.cdn}.${config.domain.domain}`, "'unsafe-inline'", "'unsafe-eval'"],
            "img-src": ["'self'", `${config.domain.subdomain.cdn}.${config.domain.domain}`],
            "default-src": ["'self'", `${config.domain.subdomain.app}.${config.domain.domain}`, `${config.domain.subdomain.api}.${config.domain.domain}`, "'unsafe-inline'", "'unsafe-eval'"]
        },
    })
);
// ====================================================== //
// ============== Add Session to front  ================= //
// ====================================================== //

app.use(function(req, res, next) {
    res.setHeader("x-powered-by", " ⭐ Powered By Magic 🦄 ");
    res.header('Access-Control-Allow-Origin', '*');
    res.header('Access-Control-Allow-Credentials', 'true');
    res.header('Access-Control-Allow-Methods', 'GET, POST');
    next();
});

// ====================================================== //
// ============== Setup connect middleware ============== //
// ====================================================== //

const apiConnect = connect();
const frontConnect = connect();
const cdnConnect = connect();

// ====================================================== //
// =============== Create route middleware ============== //
// ====================================================== //

apiConnect.use(require("./router/api.js"));
frontConnect.use(require("./router/front.js"));
cdnConnect.use(serveStatic("./public"));

// ====================================================== //
// ==== Execute express middleware (express configs) ==== //
// ====================================================== //
con.use(app);

// ====================================================== //
// ======== Execute express middleware (vhost) ========== //
// ====================================================== //

con.use(
    vhost(
        config.links.api.substring(config.links.api.lastIndexOf("/") + 1),
        apiConnect
    )
);
con.use(
    vhost(
        config.links.cdn.substring(config.links.cdn.lastIndexOf("/") + 1),
        cdnConnect
    )
);

// ====================================================== //
// ======== Execute express middleware (front)  ========= //
// ====================================================== //
con.use(frontConnect);

app.get("*", (req, res, next) => {
    next();
});

// ====================================================== //
// =================== Listen on port =================== //
// ====================================================== //
con.listen(config.system.port, () => {
    console.log(`Domain: ${config.domain.domain}`);
    console.log(`Protocol: ${config.domain.protocol}`);
    console.log(`Port: ${config.system.port}`);
    console.log('\nServer running...')
});
```

# /middleware/database/cert.pem
```
{YOUR CERT HERE...}
```

# /middleware/database/key.pem
```
{YOUR KEY HERE...}
```

# /middleware/database/index.js
```
const { MongoClient, ServerApiVersion } = require('mongodb');
const fs = require('fs');

const client = new MongoClient('mongodb+srv://##########.mongodb.net/#########?authSource=%24external&authMechanism=MONGODB-X509&retryWrites=true&w=majority', {
  sslKey: './middleware/database/key.pem',
  sslCert: './middleware/database/cert.pem',
  serverApi: ServerApiVersion.v1
});

// QueryData(database.client, "Main", "statistics", { name: "user-count" })
async function QueryData(Client,DatabaseName,Collection,Query) {
  try {
    await Client.connect();
    const database = await Client.db(`${DatabaseName}`);
    col = await database.collection(`${Collection}`)
    const data = await col.findOne(Query);
    if(!data || data == null){await Client.close();return false}else{await Client.close();return data}
  } catch(error) {
    await Client.close()
    return null
  }
}

exports.QueryData = QueryData;
exports.client = client;
```

# /config/config.js
```
// ====================================================== //
// ================ Create system object ================ //
// ====================================================== //
const system = {
    'port': '80', // Port to listen on
    'socket': '3000', // Socket port to listen on
    'version': '1.0.0', // Version of the system
    'proxy': '1', // Proxy bypass amount
    session: {
        secret: '',
        name: '',
    },
}

// ====================================================== //
// ================ Create domain object ================ //
// ====================================================== //

const domain = {
    'domain': "example.com", // Domain to listen on
    'protocol': 'https', // Protocol to listen on
    subdomain: {
        // api.domain.com
        'api': 'api',
        'api_enable': true,

        // cdn.domain.com
        'cdn': 'cdn',
        'cdn_enable': true,
    }
}

// ====================================================== //
// ================= Create seo object ================== //
// ====================================================== //
const seo = {
    'title': 'Example',
    'description': 'Example description here.',
    'favicon': 'https://cdn.example.com/image.png',
}

// ====================================================== //
// ================ Create links object ================= //
// ====================================================== //
const links = {
    'main': `${domain.protocol}://${domain.domain}`,
    'api': domain.subdomain.api_enable == true ? `${domain.protocol}://${domain.subdomain.api}.${domain.domain}` : `${domain.protocol}://${domain.domain}/api`,
    'cdn': domain.subdomain.cdn_enable == true ? `${domain.protocol}://${domain.subdomain.cdn}.${domain.domain}` : `${domain.protocol}://${domain.domain}/cdn`
};

// ====================================================== //
// ================ Create result object ================ //
// ====================================================== //
const result = {
    system,
    domain,
    seo,
    links,
}

// Export the result
module.exports = result;
```

# /router/api.js
```
const config = require("../config/config.js");
const express = require("express"),
  router = express.Router();
const bcrypt = require("bcrypt");
const rounding = 5;
const cookieParser = require("cookie-parser");
const database = require("./../middleware/database");
const cli = database.client;


router.get("*", (res,res) => {
    return res.json("value":true,"endpoint":"api");
});

// Export the router
module.exports = router;
```

# /router/front.js
```
const config = require("../config/config.js");
const express = require("express"),
  router = express.Router();
const bcrypt = require("bcrypt");
const rounding = 5;
const cookieParser = require("cookie-parser");
const database = require("./../middleware/database");
const cli = database.client;


router.get("*", (res,res) => {
    return res.json("value":true,"endpoint":"front");
});

// Export the router
module.exports = router;
```

# /views/index.ejs
```
Hello World
```

# .gitignore
```
# Logs
logs
*.log
npm-debug.log*
yarn-debug.log*
yarn-error.log*
lerna-debug.log*

# Diagnostic reports (https://nodejs.org/api/report.html)
report.[0-9]*.[0-9]*.[0-9]*.[0-9]*.json

# Runtime data
pids
*.pid
*.seed
*.pid.lock

# Directory for instrumented libs generated by jscoverage/JSCover
lib-cov

# Coverage directory used by tools like istanbul
coverage
*.lcov

# nyc test coverage
.nyc_output

# Grunt intermediate storage (https://gruntjs.com/creating-plugins#storing-task-files)
.grunt

# Bower dependency directory (https://bower.io/)
bower_components

# node-waf configuration
.lock-wscript

# Compiled binary addons (https://nodejs.org/api/addons.html)
build/Release

# Dependency directories
node_modules/
jspm_packages/

# TypeScript v1 declaration files
typings/

# TypeScript cache
*.tsbuildinfo

# Optional npm cache directory
.npm

# Optional eslint cache
.eslintcache

# Microbundle cache
.rpt2_cache/
.rts2_cache_cjs/
.rts2_cache_es/
.rts2_cache_umd/

# Optional REPL history
.node_repl_history

# Output of 'npm pack'
*.tgz

# Yarn Integrity file
.yarn-integrity

# dotenv environment variables file
.env
.env.test

# parcel-bundler cache (https://parceljs.org/)
.cache

# Next.js build output
.next

# Nuxt.js build / generate output
.nuxt
dist

# Gatsby files
.cache/
# Comment in the public line in if your project uses Gatsby and *not* Next.js
# https://nextjs.org/blog/next-9-1#public-directory-support
# public

# vuepress build output
.vuepress/dist

# Serverless directories
.serverless/

# FuseBox cache
.fusebox/

# DynamoDB Local files
.dynamodb/

# TernJS port file
.tern-port
```

command name | command
--- | ---
Setup npm package manager | npm init 
Install vhost package| npm i vhost 
Install connect package| npm i connect 
Install express package| npm i express 
Install express-minify-html package| npm i express-minify-html 
Install serve-static package| npm i serve-static 
Install compression package| npm i compression 
Install cookie-parser package| npm i cookie-parser 
Install express-session package| npm i express-session 
Install mongodb package| npm i mongodb
