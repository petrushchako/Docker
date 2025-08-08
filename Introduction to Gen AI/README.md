# Introduction to Gen AI

## Docker Model Runner
Run GenAI Model Locally

### Key Features
- **Simple Local Model Execution**: based on llama.cpp, OpenAI API compliant
- **GPU Acceleration**:
  - At the beginning: It offers GPU acceleration on Apple silicon
  - Now: Windows+Nvidia
- **Standardized Model Packaging**: Open Container Initiative (OCI) Artifacts

### Future Plans
- **More platforms**: Linux, Docker CE
- **Ability to customize and publish your own models**
- **Better integration with development workflows** (including Compose and Testcontainers)

<br><br><br>

## Docker Model Runner (`docker model`)
- Avaialble with Docker Desktop 4.40
- Make sure Docker Destop is running prior to execution of any commands below.
- Use `docker model list` command to view the list of downloaded OCI artifacts available locally
- Pull ai models with `docker model pull <image_name>` eg. `docker model pull ai/llama3.2`
- Run docker AI model with `docker model run ai/llama3.2`
- Quick chat with `/bye`
