pipeline {
    agent any
    
    tools {
      go 'go'
    }


    environment {
        GITHUB_TOKEN = credentials('Github')
    }

    stages {
        stage('Clone Repository') {
            steps {
                sh('rm -rf devops-homework-1 && git clone https://${GITHUB_TOKEN}@github.com/brayandm/devops-homework-1.git')
            }
        }
    }
}
