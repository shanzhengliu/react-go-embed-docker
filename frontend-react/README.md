
### Introduction  
This file is used to create embed project with React(frontend) and Go(backend).  

### React  
This project tech stack are  
1. vite `build` 
2. css `tailwindCSS`  
3. swr `data fetching hooks`  
4. backend mock `mockjs`  
5. data fectching library: `axios`  
you need to run commend `cd frontend-react` to step into frontend project and run the command.
## command list  
1. `pnpm install`: dependency install
2. `pnpm dev`: api will access real backend. and updated the value in .env.developemtn, default is `http://localhost:8080`, which is local backend project.  
3. `pnpm mock`: use the mock data as backend. you can create new mock in folder `mock`.  
4. `pnpm backend`, will run the go backend appliction locally.

### backend project.
This is the project for create api. at the moment is pure and we are 
using `github.com/gorilla/mux` as http web client.  

1. the defualt cross origin is `*`, and you can set up in the environment variable `ORIGIN`  
2. the default port is `8080`, you can set up in environment variable `PORT`

### Docker build.
You can easily build the project including frontend and backend via command `Docker build .t `