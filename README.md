# Go GraphQL Mongo Service

This Go project demonstrates the usage of GraphQL to create job listings. Users can create, update, and delete these job listings using various GraphQL queries and mutations. The project utilizes the [gqlgen](https://github.com/99designs/gqlgen) library to simplify GraphQL server implementation.

## Installation

1. **Clone the repository:**

```bash
git clone --recursive https://github.com/pi-prakhar/go-graphql-mongo.git
git submodule update --init --recursive
```

3. **To start the project in Docker:**

```bash
docker-compose up -d --build
```
You can access the server at `localhost:8000`.

4. **To start the project locally:**

*To run the project locally, ensure you have MongoDB running on your machine. If you are using MongoDB Atlas, refer to the "Configure MongoDB Atlas" section below.
*Change the configuration in `config/config.json` by setting the `docker` property to `false`.
*After this, run commands [Windows]:

```bash
go build -o main.exe ./cmd/go-graphql-mongo/main.go
.\main.exe
```

You can access the server at `localhost:8000`.

## Usage

To access the GraphQL playground, go to `localhost:8000/playground`.

The GraphQL query endpoint is `localhost:8000/api/query`.

### Queries

#### Get All Jobs

```graphql
query GetAllJobs {
  jobs {
    _id
    title
    description
    company
    url
  }
}
```

#### Get Job By Id

```graphql
query GetJob($id: ID!) {
  job(id: $id) {
    _id
    title
    description
    url
    company
  }
}
```

Variables:
```json
{
  "id": ""
}
```

### Mutations

#### Create Job

```graphql
mutation CreateJobListing($input: CreateJobListingInput!) {
  createJobListing(input: $input) {
    _id
    title
    description
    company
    url
  }
}
```

Variables:
```json
{
  "input": {
    "title": "",
    "description": "",
    "company": "",
    "url": ""
  }
}
```

#### Update Job By Id

```graphql
mutation UpdateJob($id: ID!, $input: UpdateJobListingInput!) {
  updateJobListing(id: $id, input: $input) {
    title
    description
    _id
    company
    url
  }
}
```

Variables:
```json
{
  "id": "",
  "input": {
    "title": "",
    "description": "",
    "company": "",
    "url": ""
  }
}
```

#### Delete Job By Id

```graphql
mutation DeleteQuery($id: ID!) {
  deleteJobListing(id: $id) {
    deletedJobId
  }
}
```

Variables:
```json
{
  "id": ""
}
```

## Interact with the Service using client :

To interact with the service in interactive mode, you can use the web client built using Next.js.

To start the client:

1. CD to `/client-web/nextjs-graphql`.
2. Create .env.local similar to .env.sample.
3. Make sure backend server is accessible at localhost:8000 in local machine.
4. Run the command:
   ```bash
   npm i
   npm run dev
   ```
5. Access the project at `localhost:3000`.

## Configure MongoDB Atlas

You can configure MongoDB Atlas for the current project if you do not wish to run the project in Docker or use local MongoDB Compass.

Change the `mongo-atlas` property to `true` in `config/config.json`. Create a `.env` file similar to `.env.sample` and add your MongoDB URI link.

---

Let me know if you need further modifications!
