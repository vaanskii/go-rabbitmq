# RabbitMQ Go Producer and Consumer

This project demonstrates a `Go` Producer and Consumer using `RabbitMQ.`


- `Producer`: Sends messages to a `RabbitMQ` queue via an `HTTP API` built with `Gin`.
- `Consumer`: Consumes messages from the queue and logs them.


### How it works
- The `Producer` exposes an endpoint GET `/send?msg=message`to send a message to the queue.
-  The `Consumer` listens to the queue and logs any incoming messages.


### Setup
- Install dependencies:
```bash
  go mod tidy
```

  - Run the Producer and Consumer.
