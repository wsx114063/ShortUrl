pipeline {
    agent any

    stages {
        stage('Build Golang Environment') {
            steps {
                script {
                    checkout scm

                    def dockerImage = docker.build('shortenurl', '-f Dockerfile .')
                    dockerImage.run('-d -p 8081:8081')
                    // dockerImage.inside {
                    //     sh 'go build -o my-app'
                    // }
                    // stash includes: 'my-app', name: 'build-artifacts'
                }
            }
        }
        // stage('Test') {
        //     steps {
        //         sh 'ls -l'

        //         unstash 'build-artifacts'
                
        //         sh 'ls -l'
        //     }
        // }
        // stage('Deployee') {
        //     steps {
        //         script {
        //             def dockerRun = docker.image('shortenurl').run('-d -p 8081:8081')
        //         }
        //     }
        // }
    }
}
