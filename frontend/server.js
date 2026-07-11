import http from "node:http";
import { createReadStream, promises as fs } from "node:fs";
import path from "node:path";
import { fileURLToPath } from "node:url";

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

const MIME_TYPES = {
  ".html": "text/html; charset=utf-8",
  ".css": "text/css; charset=utf-8",
  ".js": "application/javascript; charset=utf-8",
  ".json": "application/json; charset=utf-8",
  ".png": "image/png",
  ".jpg": "image/jpeg",
  ".jpeg": "image/jpeg",
  ".gif": "image/gif",
  ".svg": "image/svg+xml",
  ".ico": "image/x-icon",
  ".woff": "font/woff",
  ".woff2": "font/woff2",
};

const server = http.createServer(async (req, res) => {
  try {
    const url = new URL(req.url, `http://${req.headers.host}`);

    let filePath = path.join(
      __dirname,
      url.pathname === "/" ? "index.html" : url.pathname
    );

    // Prevent path traversal
    filePath = path.normalize(filePath);

    if (!filePath.startsWith(__dirname)) {
      res.writeHead(403);
      return res.end("Forbidden");
    }

    try {
      const stat = await fs.stat(filePath);

      if (stat.isFile()) {
        const ext = path.extname(filePath).toLowerCase();
        const type = MIME_TYPES[ext] || "application/octet-stream";

        res.writeHead(200, { "Content-Type": type });
        return createReadStream(filePath).pipe(res);
      }
    } catch {}

    // SPA fallback
    const indexPath = path.join(__dirname, "index.html");
    const html = await fs.readFile(indexPath);

    res.writeHead(200, { "Content-Type": "text/html; charset=utf-8" });
    res.end(html);
  } catch (err) {
    res.writeHead(500);
    res.end("Internal Server Error");
  }
});

server.listen(3000, () => {
  console.log("🚀 Server running at http://localhost:3000");
});
