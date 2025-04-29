pipeline {
    agent { docker { image 'golang:1.24.2-alpine3.21' } }
    stages {
        stage('build') {
            steps {
                sh 'GOOS=linux GOARCH=amd64 go build -o build ./cmd/app'
            }
        }
    }
}
