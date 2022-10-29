1. Init frontend: 
```sh
$ npm create vite@latest front -- --template vue-ts
````
2. Install dependencies:
```sh
$ cd front && npm i
````
3. Build:
```sh
cd front && npx vue-tsc && npx vite build
```

front/index.html can be edited.

front/dist will be embedded in server binary and served under "/" endpoint.
