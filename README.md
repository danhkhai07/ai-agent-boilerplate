# ai-agent-app

## How to run

1. Install Docker
2. Create an `.env` file with these variables:

    ```env
    PORT=8080
    AGENT_STATIC_DIR=/path/to/static/files
    GEMINI_API_KEY=...  # valid Gemini token. Create at Google AI Studio
    ```

3. Run the following commands:

    ```bash
    sudo docker build -t agent-app .
    sudo docker run -d -p 8080:8080 --env-file /path/to/.env agent-app
    ```

4. The app is hosted at port 8080. Open your browser at `http://localhost:8080`
