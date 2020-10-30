pipeline {
    agent any
    tools {
        go 'go'
        dockerTool 'docker'
    }
    stages {
        stage ('Install dependencies') {
            steps {
                sh 'go get github.com/gorilla/mux'
                sh 'go get github.com/go-sql-driver/mysql'
                sh 'go get github.com/joho/godotenv'
            }
        }
        stage ('Git') {
            steps {
                git url: 'https://github.com/emorales8/Registration-service'
            }
        }
        stage ('Go App Build') {
            steps {
                sh 'go build post.go'
            }
        }
        stage ('Docker build') {
            steps {
                sh 'docker build -t go-post-srv:v1.0 .'
            }
        }
    }
}
