# Go High-Performance Distributed Task Scheduler & Reminder Daemon

A lightweight, concurrent, time-driven background service built natively in Go to poll cloud data streams and trigger precision automated workflows. This service replaces resource-heavy cron intervals by utilizing system-hardware-driven tickers to process high-volume asynchronous event queues with zero framework dependencies.

---

### 🛠️ Production Architecture

*   **Time-Driven Heartbeat Engine:** Built using Go's native `time.Ticker` channel primitives (`ticker.C`) to establish precise execution heartbeats driven directly by the CPU's hardware clock layer.
*   **Asynchronous Processing Pipeline:** Leverages concurrent `goroutines` to spawn isolated, thread-safe worker instances for each due task. This allows thousands of simultaneous alerts to process in parallel without blocking the main time synchronization loop.
*   **Atomic Data Synchronization:** Integrated with standard `pgxpool` high-performance connection structures to query and mutate task states inside a cloud **Neon Serverless Postgres** database cluster.
*   **Decoupled State Mutation Engine:** Executes atomic `UPDATE` operations immediately following task verification to transition records to a completed state (`is_completed = TRUE`), preventing duplicate execution and race conditions across the worker pool.
*   **Environment Configuration Isolation:** Utilizes automated environment variable loaders (`godotenv`) combined with `.gitignore` parameters to cleanly inject infrastructure connection string secrets without exposing sensitive cloud parameters to source control.

---

### 📊 Verified Live Execution Logs

During live execution loops connected directly to a remote Neon Postgres database cluster, the scheduler engine captures, scales, and shuts down batches smoothly:

```text
⏰ Task Scheduler Engine Initializing...
Successfully connected to Neon serverless Postgres Cluster!
🚀 Time-Driven Engine Online. Scanning for pending records.

⏱️ [Heartbeat Clock tick]: 11:56:32
🔥 Found 2 due tasks! Processing batch payload...
[Job #1] Executing action targeted for recipient: 628123456789
[Job #1] Dispatching content stream: "🚨 [ALERT]: System health check verification reminder from Neon!"
[Job #2] Executing action targeted for receptionist: 628987654321
[Job #2] Dispatching content stream: "💳 [FINANCE]: Inbound subscription renewal invoice notice from Neon!"
✅ [Job #1] State permanently committed as COMPLETED in cloud storage.
✅ [Job #2] State permanently committed as COMPLETED in cloud storage.

⏱️ [Heartbeat Clock tick]: 11:56:42
😴 No tasks due in this cycle. Waiting for next tick...
```
