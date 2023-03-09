package com.tdc.demo.infrastructure.adapter.output;

import java.time.LocalDateTime;

import lombok.AllArgsConstructor;
import lombok.NoArgsConstructor;
import org.springframework.amqp.core.AmqpTemplate;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.scheduling.annotation.Scheduled;
import org.springframework.stereotype.Component;

import com.tdc.demo.domain.boundary.output.Producer;
import com.tdc.demo.domain.dto.MessageDTO;

@Component
@AllArgsConstructor
@NoArgsConstructor
public class RabbitMQProducer implements Producer {

    @Value("${rabbitmq.queue}")
    private String queueName;

    @Autowired
    private AmqpTemplate template;

    private LocalDateTime timestamp;
    @Scheduled(fixedDelay = 1000)
    public void publish() {
        timestamp = LocalDateTime.now();
        MessageDTO messageDTO = new MessageDTO("Hello world", timestamp);
        this.template.convertAndSend(queueName, messageDTO);
    }
}
