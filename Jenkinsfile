pipeline {
    agent any

    stages {
        stage('Build Golang Environment') {
            steps {
                script {
                    checkout scm
                    sh 'docker build -t shortenurl .'
                    stash includes: '.', name: 'build-artifacts'                   
                }
            }
        }
        stage('Test') {
            steps {
                sh 'pwd'
                sh 'ls -l'
                unstash 'build-artifacts'                
                sh 'ls -l'
            }
        }
        stage('Deployee') {
            steps {
                sh 'docker run -d -p 8081:8081 --name shortenurl shortenurl'
            }
        }
    }
}
