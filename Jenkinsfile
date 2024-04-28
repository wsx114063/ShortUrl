pipeline {
    agent any

    environment {
        GOROOT = '/usr/local/go'
        PATH = "/usr/local/go/bin:${env.PATH}"
    }

    stages {
        stage('Build Golang Environment') {
            steps {
                script {
                    checkout scm
                    sh 'chmod +x build.sh'
                    sh './build.sh'
                }
            }
        }
        stage('Test') {
            steps {
                script {
                    def testOutput = sh(script: 'go test ./...', returnStdout: true).trim()
                    if (testOutput.contains('PASS')) {
                        echo 'Tests passed, proceeding with deployment'
                    } else {
                        error 'Tests failed, deployment aborted'
                    }
                }
            }
        }
        stage('Deployee') {
            steps {
                sh 'docker-compose up -d -p shortenurl --build'
            }
        }
    }
    post {
        always {
            discordSend description: "Build:${currentBuild.number} Status: ${currentBuild.currentResult}", footer: "Footer Text", link: "${env.BUILD_URL}", result: "${currentBuild.currentResult}", title: "${JOB_NAME}", webhookURL: "https://discord.com/api/webhooks/1232882729343778826/O5asexU5APt5XlUeoevg-hc7lB9xAjuVYaHjdYE-awSnBmXr1jvj4DaWffOEJJvizwN5"
        }
    }
}
