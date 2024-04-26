pipeline {
    agent { dockerfile true }

    stages {
        stage('Build Dockerfile') {
            steps{
                echo 'test build dockerfile'
            }
            post {
                always {
                    discordSend description: "Build:${currentBuild.number} Status: ${currentBuild.currentResult}", footer: "Footer Text", link: "${env.BUILD_URL}", result: "${currentBuild.currentResult}", title: "${JOB_NAME}", webhookURL: "https://discord.com/api/webhooks/1232882729343778826/O5asexU5APt5XlUeoevg-hc7lB9xAjuVYaHjdYE-awSnBmXr1jvj4DaWffOEJJvizwN5"
                }
            }
        }
    }    
}