pipeline {
    agent any

    stages {
        stage('Build Golang Environment') {
            steps {
                script {
                    checkout scm
                    sh 'pwd'
                    sh 'ls -l'
                    sh 'docker build shortenurl .'

                    // def dockerImage = docker.build('shortenurl', '-f Dockerfile .')
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
