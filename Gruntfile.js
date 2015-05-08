module.exports = function (grunt) {

    // Project configuration.
    grunt.initConfig({
        pkg: grunt.file.readJSON('package.json'),
        config: grunt.file.readJSON('config.json'),
        bintrayDeploy: {
            bintray: {
                options: {
                    user: "<%= config.bintray.user %>",
                    apikey: "<%= config.bintray.apikey %>",
                    pkg: {
                        repo: "generic"
                    }
                },
                files: [{
                    src: ["build/*"],
                    dest: "<%= pkg.version %>",
                    filter: "isFile"
                }]
            }
        },
        "vagrant_commands": {
            default: {
                commands: [
                    ['halt'],
                    ['up', '--provision']
                ]
            }
        },
        go: {
            "terraform-couchdb": {
                root: "src",
                output: "../build/terraform-provider-couchdb"
            }
        }
    });
    grunt.loadNpmTasks('grunt-go');
    grunt.loadNpmTasks('grunt-vagrant-commands');
    grunt.loadNpmTasks('grunt-bintray-deploy');

    // Default task(s).
    grunt.registerTask('default', ['vagrant_commands', 'go:build:terraform-couchdb']);

};