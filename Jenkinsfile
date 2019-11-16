pipeline {
    agent {
        docker { image 'golang:1.13-alpine' }
    }
    stages {
        stage('build') {
            steps {
                sh 'go version'
            }
        }
    }
}
