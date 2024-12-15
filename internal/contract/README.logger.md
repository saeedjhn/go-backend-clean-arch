### Here’s a list of all the suggested logging formats, categorized by type and complexity

1. **Basic Format**

   Structure:
    ```
   <Operation>.<Details> - Additional description
   
   User.Login - User logged in successfully  
   Migration.Up - Applying migrations for version 1.2.3
   ```

2. **Service-Specific Format**

   Structure:
   ```
   <ServiceName>.<Operation>.<Details> - Additional description
   
   UserService.User.Login - User logged in successfully  
   OrderService.Order.Create - Order created, order_id: 98765
   ```

3. **Other**
```
Infof("Queue bound to exchange, starting to consume from queue, consumerTag: %v", consumerTag)
Errorf("Failed to process delivery: %v", err)
Errorf("Failed to acknowledge delivery: %v", err)
Info("Deliveries channel closed")
Errorf("ch.NotifyClose: %v", chanErr)
Infof("AppVersion: %s, LogLevel: %s, Mode: %s, SSL: %v", cfg.Server.AppVersion, cfg.Logger.Level, cfg.Server.Mode, cfg.Server.SSL)
Fatalf("Postgresql init: %s", err)
Infof("PostgreSQL connected: %#v", psqlDB.Stats())
Errorf("signal.Notify: %v", v)
Errorf("ctx.Done: %v", done)
Errorf("Metrics router.Shutdown: %v", err)
Info("Server Exited Properly")
```