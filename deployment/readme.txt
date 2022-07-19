1. Podesiti dokerfile za pokretanja golanga (u folderu servers) sa portom kao argumentom (ova izmenio sam ja da golang cita parametar PORT)

2. Pozicionirati se u terraform folderu pokrenuti: docker compose up init (to ce kreirati app na heroku)

3. Initiovati terraform komanda: terraform init   (choco install terraform   - komanda za instaliranej terraforma na win )
(komanda za instaliranje choco iz powerShell-a - pokrenuti kao administrator:
" Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1')) ")

4. Prilikom intalacije trazice da se postgresql connection string
(Database Credentials - https://data.heroku.com/datastores/2e6d740e-417a-4dd4-8d41-cd78ad082823#administration  mislim da se odatle preuzimaju kredencijali,
proveriti da li format odgovara agentskoj aplikaciji lokalno)

5. Pokrenuti: docker compose up deploy

6. to je to

proveriti da li se kreirala aplikacija devops22tim8 na https://id.heroku.com/login
mail:damjan0032@gmail.com
pass:DevOps123.

Link vezbi: https://www.youtube.com/watch?v=9nyOwEkIpv4&list=PLnZTWZfntZ9AriSSsO6KnkGP4-qeDecFP&index=3&ab_channel=DevOps