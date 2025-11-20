# Sovereign UI

A sleek, modern web interface for managing your Headscale network.

## Phase 1: MVP

This initial version provides a read-only dashboard to view your Headscale users and their connected nodes.

### Features

- **Dark Theme:** A modern, easy-on-the-eyes interface.
- **User/Node Listing:** View all users and their associated devices.
- **Status Indicators:** See which nodes are currently online.
- **Responsive Design:** Works on desktop and mobile.

### Quick Start with Docker Compose

**Prerequisites:**
- Docker and Docker Compose installed.
- A running Headscale instance.
- A Headscale API key.

**1. Generate a Headscale API Key:**
On your Headscale server, run:
```bash
headscale apikey create
```
Copy the key.

**2. Configure Sovereign UI:**
In the root of this project, create a `.env` file for the backend:
```bash
cp backend/.env.example backend/.env
```
Edit the newly created `backend/.env` file and add your Headscale API URL and the API key you just generated.
```
HEADSCALE_API_URL=http://your-headscale-ip:8080
HEADSCALE_API_KEY=your-generated-api-key
```
**Note:** If your Headscale instance is running on the same machine as Docker, you can use `http://host.docker.internal:8080` as the API URL.

**3. Run the Application:**
From the root of the `sovereign-ui` directory, run:
```bash
docker-compose up --build
```

**4. Access the UI:**
Open your web browser and navigate to `http://localhost`.

You should now see the Sovereign UI dashboard, displaying the users and nodes from your Headscale network!