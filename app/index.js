'use strict'

const path = require('path')
const Generator = require('yeoman-generator')
const mkdir = require('mkdirp')

module.exports = class extends Generator {
    constructor(args, opts) {
        super(args, opts)

        this.argument("appname", {type: String, required: true})
        this.option("appdir", {type: String, default: "."})
    }


    writing() {
        console.log('\n' +
            'ISV CI idiomatic Go CLI app\n' +
            '---------------------------\n' +
            '\n')

        const appName = this.options.appname.replace(/\s+/g, '-').toLowerCase()
        const location = this.options.appdir + ((this.options.appdir.substr(-1) === '/') ? '' : '/')
        const appDir = location + appName

        console.log('Generating app dir \'' + appDir + '\'')

        // let pkgDir = this.destinationPath('pkg');
        // let srcDir = this.destinationPath(path.join('src/', this.repoUrl));
        // let binDir = this.destinationPath('bin');
        //
        // mkdir.sync(pkgDir);
        // mkdir.sync(srcDir);
        // mkdir.sync(binDir);
        //
        // this.fs.copy(
        //     this.templatePath('_gitignore'),
        //     path.join(srcDir, '.gitignore')
        // );
        // this.fs.copy(
        //     this.templatePath('_hello.go'),
        //     path.join(srcDir, '/hello/hello.go')
        // );
        // this.fs.copy(
        //     this.templatePath('_hello_test.go'),
        //     path.join(srcDir, '/hello/hello_test.go')
        // );
        //
        // let tmplContext = {
        //     appName: this.appName,
        //     repoUrl: this.repoUrl
        // };
        //
        // this.fs.copyTpl(
        //     this.templatePath('_main.go'),
        //     path.join(srcDir, 'main.go'),
        //     tmplContext
        // );
        // this.fs.copyTpl(
        //     this.templatePath('_README.md'),
        //     path.join(srcDir, 'README.md'),
        //     tmplContext
        // );
        // this.fs.copyTpl(
        //     this.templatePath('_Makefile'),
        //     path.join(srcDir, 'Makefile'),
        //     tmplContext
        // );

    }
}
