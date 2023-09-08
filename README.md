# TTRPG Compass
Web-based tool to view and compare various TTRPG systems (currently uses generated mock data). 

Key features
------------
- **Aspect stellar chart** ~ quickly view what the focus of a given system is across different aspects of play (such as _Complexity_, _Combat_, _Exploration_ etc.)
- **General information views** ~ _Name_, _Genre_, _Family_, _GM/GMless_ etc.
- **Similarity view** ~ display 3 most similar systems based on aspects, genre and family similarity

Development
-----------
Some of the important info about each of the 4 separate parts of the app.

**I. Client** (SvelteKit)

Should have an `.env` file with the `API_URL` specified. Uses `pnpm` for package management - either use `pnpm run dev` or spin up the Dockerfile. Uses SvelteKit with the `node adapter`.

**II. API** (Go + Fiber)

Either spin up the Dockerfile, or setup `go`. Connects to DB and rebuilds it on startup from the provided folder (`e.g. data/mock/`). Expects an `.env` file with the following: `CLIENT_URL`, `REDIS_URL`, and `REDIS_PASS`. 

**III. Database** (Redis Stack)

You can spin up Redis stack instance for dev purposes with the recommended `docker run -d --name redis-stack -p 6379:6379 -p 8001:8001 redis/redis-stack:latest` or create a free instance on Redis Cloud. 

An index needs to be manually added with the command `FT.CREATE idx:systems ON HASH PREFIX 1 system: SCHEMA title TEXT NOSTEM genre TEXT family TEXT`.

**IV. Mock data generation toolkit** (Python)

You can generate mock systems with `python3 init_systems.py N` where N is number of systems to generate (default folder is `../mock`). 

You can delete the mock data with `python3 delete_systems.py`, if you provide the `-f` flag it will also flush the local redis db instance (`docker exec -it redis-stack redis-cli FLUSHALL`).

Similarity weights for genre and family are in the `calc_similarity` function. 
