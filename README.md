# HTTP server dockerized to convert files to PDF
HTTP server written in Go (golang) that uses unoconv to transform any file to PDF.

To start the server:
1. Download the repo `git clone https://github.com/arturo-source/pdf-converter-docker.git`
2. Build the image `docker build -t pdf-converter .`
3. Run the container `docker run -d -p 8000:8000 pdf-converter`

To use the service you have to do a request to `/convert-to-pdf` endpoint , parsing the file with multipart/form-data. Example with curl: `curl -F "file=@source.docx" localhost:8000/convert-to-pdf > target.pdf`.

The point of this repository is to make your own microservice, but it is fully functional.
But it bothers me that final image occupies 1.94GB. If you can help me building a lighter image changing the Dockerfile, **I need your pull request**.
