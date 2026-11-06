# Go High-Performance Distributed Task Scheduler & Outbound Engine

A lightweight, concurrent, time-driven background service built natively in Go to poll cloud data streams, manage structured data serialization, and dispatch automated outbound transactional workflows. This service coordinates system-hardware-driven tickers to process high-volume asynchronous event queues with zero framework dependencies.

---

### 🛠️ Production Architecture

*   **Time-Driven Heartbeat Engine:** Built using Go's native `time.Ticker` channel primitives (`ticker.C`) to establish precise execution heartbeats driven directly by the CPU's hardware clock layer.
*   **Asynchronous Processing Pipeline:** Leverages concurrent `goroutines` to spawn isolated, thread-safe worker instances for each due task. This allows thousands of simultaneous alerts to process in parallel without blocking the main time synchronization loop.
*   **Modular Multi-File Packaging:** Split into decoupled packages (`main.go`, `scheduler/models.go`, `scheduler/worker.go`) to isolate configuration, data schema blueprints, and network transaction runtimes cleanly.
*   **Outbound Network Integration:** Incorporates a native Go `http.Client` block with strict timeout boundary safety thresholds to serialize multi-parameter JSON data payloads directly to third-party endpoints (Meta Graph API).
*   **Atomic Data Synchronization:** Integrated with standard `pgxpool` high-performance connection structures to query and mutate task states inside a cloud **Neon Serverless Postgres** database cluster immediately following validation to prevent race conditions.

---

### 📊 Verified Live Execution Logs

During live execution loops connected directly to a remote Neon Postgres database cluster, the scheduler engine tracks ticks, queries tasks, and dispatches JSON network payloads smoothly:

```text
⏰ Task Scheduler Engine Initializing...
Successfully connected to Neon serverless Postgres Cluster!
🚀 Time-Driven Engine Online. Scanning for pending database records.

⏱️ [Heartbeat Clock tick]: 21:01:23
🔥 Found 1 due tasks! Processing batch payload...
[Job #3] Dispatching execution sequence to: 6281249138361
[Outbound Success]: Alert successfully dispatched across the internet to 6281249138361
✅ [Job #3] State permanently committed as COMPLETED in cloud storage.

⏱️ [Heartbeat Clock tick]: 21:01:33
😴 No tasks due in this cycle. Waiting for next tick...
```
