Downloads resumes automatically if they had been interrupted—no need for manual restart. Resumes from where it left off, thanks to HTTP Range support.

Supports downloading of multiple files simultaneously, leveraging Go's goroutines and sync primitives for safe and efficient concurrency that won't tie up resources.

Intelligently determines filenames by looking at the server headers (such as Content-Disposition) or resorting to the filename in the URL itself.

Retries failed downloads elegantly, by using HTTP status codes to determine whether to retry, and delaying retries with a configurable timeout.

Displays progress bars per file so you can monitor download speeds and completion in real-time, directly in your terminal.

Avoids already completed files, so reruns are safe, bandwidth-efficient, and quicker—no duplicate work.

Kept in clean, testable modules, ensuring the codebase is simple to extend, modify, or transform into other tools.
