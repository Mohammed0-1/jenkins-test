pipeline {
    agent { docker { image 'golang:1.24.2-alpine3.21' } }
    
    stages {
                stage('cleanup') {
            steps {
                // Clean up the workspace before starting the build
                sh 'rm -rf build '
            }
        }

        stage('build') {
            steps {
                // Build the Go executable
                sh 'GOOS=linux GOARCH=amd64 go build -o build/myapp ./cmd/app'
            }
        }
        
        stage('archive') {
            steps {
                // Archive the built executable as a Jenkins artifact
                archiveArtifacts artifacts: 'build/myapp', allowEmptyArchive: true
            }
        }
    }
}
