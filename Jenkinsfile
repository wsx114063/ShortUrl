pipeline {
    agent any

    stages {
        stage('Build Golang Environment') {
            steps {
                script {
                    checkout scm
                    sh 'go build .'
                    sh 'ls -l'
                }
            }
        }
        stage('Test') {
            steps {
                sh 'pwd'               
                sh 'ls -l'
            }
        }
        stage('Deployee') {
            steps {
                sh 'docker-compose up -d '
            }
        }
    }
    post {
        always {
            discordSend description: "Build:${currentBuild.number} Status: ${currentBuild.currentResult}", footer: "Footer Text", link: "${env.BUILD_URL}", result: "${currentBuild.currentResult}", title: "${JOB_NAME}", webhookURL: "https://discord.com/api/webhooks/1232882729343778826/O5asexU5APt5XlUeoevg-hc7lB9xAjuVYaHjdYE-awSnBmXr1jvj4DaWffOEJJvizwN5"
        }
    }
}
