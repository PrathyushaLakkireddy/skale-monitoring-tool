### Configure the following variables in `config.toml`
- **[skale_endpoint]**
   - *skale_node_ip*
    
      Skale node ip is used to gather information like hardware info, schains status, container health, block number etc.

- **[enable_alerts]**

   - *enable_telegram_alerts*

      Configure **yes** if you wish to get telegram alerts otherwise make it **no**.

   - *enable_email_alerts*

      Configure **yes** if you wish to get email alerts otherwise make it **no**.

- **[alerter_preferences]**
    
    - *container_health_alerts*

      Configure **yes** if you wish to get alerts when container state is *Paused* or *Dead* , otherwise configure it as **no**

    - *ethbalance_change_alerts*

      If you want to recieve alerts when the `ETH` account balance has dropped below to configured threshold make it as **yes** otherwise **no**
    
    - *sklbalance_change_alerts*
    
      If you want to recieve alerts when the `SKALE` account balance has dropped below to configured threshold make it as **yes** otherwise **no**

    - *sgx_status_alerts*
      
      If you want to recieve alerts when SGX wallet is not connected to server make it as **yes** otherwise **no**

    - *min_ETHbalance_alerts*

      If you wish to recieve alerts when ETH wallet balance has dropped below or equal to 1.5ETH, make it as **yes** otherwise **no**

    - *node_health_alert*
     
      If you wish to recieve alerts when node health is not OK, then configured it as **yes** otherwise **no**

- **[alerting_thresholds]**

   - *ethbalance_change_threshold*

      An integer value to recieve ETH balance change alerts, e.g. if your account balance has dropped to given threshold value you will receive alerts.

   - *sklbalance_change_threshold*
     
      An integer value to recieve Skale balance change alerts, e.g. if your account balance has dropped to given threshold value you will receive alerts.
    
- **[regular_status_alerts]**
  - *alert_timings*
     
     Array of timestamps for to get alerts about Container status and Node status. You will get alerts based on time which you configured.

- **[telegram]**
  - *tg_chat_id*

      Telegram chat ID to receive alerts to your telegram chat, required for Telegram alerting.
    
  - *tg_bot_token*

      Telegram bot token, required for Telegram alerting. The bot should be added to the chat and should have send message permission.
    
- **[Email]**

  - *email_address*

      E-mail address to receive mail notifications, required for e-mail alerting.
   
  - *sendgrid_token*

      Sendgrid mail service api token, required for e-mail alerting.

- **[prometheus]**

    - *prometheus_address*

      Prometheus address to export solana metrics and serve, by default listening address configured as (http://localhost:5000) in `config.toml` .

    - *listen_address*
       
      Port in which prometheus server will run,and export metrics on this port, (ex: http://localhost:5000/metrics) shows all the metrics which are stored in prometheus database, by default it will run on 9090 port.

     

    

     
      