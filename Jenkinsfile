pipeline {
    agent none

    stages {
        stage('Build') {
            agent { label 'golang' }
            steps {
                echo 'Building..'
                script {
                    checkout([$class: 'GitSCM', branches: [[name: 'CICD/test']], 
                              doGenerateSubmoduleConfigurations: false, 
                              extensions: [], submoduleCfg: [], userRemoteConfigs: [[url: 'https://github.com/wsx114063/ShortUrl.git']]])
                }
                sh 'go build -o shorten_url.exe Main.go  '
            } 
            post {
                always {
                    lastChanges format:'SIDE', matching: 'LINE'
                }
                success { 
                    archiveArtifacts artifacts: 'shorten_url.exe' 
                }
                failure {
                    discordSend description: "Build:${currentBuild.number} Status: ${currentBuild.currentResult}", footer: "Footer Text", link: "${env.BUILD_URL}", result: "${currentBuild.currentResult}", title: "${JOB_NAME}", webhookURL: "https://discord.com/api/webhooks/1232882729343778826/O5asexU5APt5XlUeoevg-hc7lB9xAjuVYaHjdYE-awSnBmXr1jvj4DaWffOEJJvizwN5"
                }
            } 
        }
        stage('Test') {
            agent { label 'test' }
            steps {
                echo 'Testing..'
                echo "download: ${BUILD_URL}/artifact/shorten_url.exe"
                sh """#!/bin/bash
                curl -O ${BUILD_URL}//artifact/shorten_url.exe
                chmod +x shorten_url.exe
                """
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying....'
            }
        }
    }
    post {
        always{
            discordSend description: "Build:${currentBuild.number} Status: ${currentBuild.currentResult}", footer: "Footer Text", link: "${env.BUILD_URL}", result: "${currentBuild.currentResult}", title: "${JOB_NAME}", webhookURL: "https://discord.com/api/webhooks/1232882729343778826/O5asexU5APt5XlUeoevg-hc7lB9xAjuVYaHjdYE-awSnBmXr1jvj4DaWffOEJJvizwN5"
        }        
    }
}