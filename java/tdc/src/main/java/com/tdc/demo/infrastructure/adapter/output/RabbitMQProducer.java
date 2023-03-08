package com.tdc.demo.infrastructure.adapter.output;

import lombok.AllArgsConstructor;
import lombok.NoArgsConstructor;
import org.springframework.amqp.core.AmqpTemplate;
import org.springframework.amqp.core.Queue;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.scheduling.annotation.Scheduled;
import org.springframework.stereotype.Component;

import com.tdc.demo.domain.boundary.output.Producer;
import com.tdc.demo.domain.dto.MessageDTO;

import java.time.LocalDateTime;

@Component
@AllArgsConstructor
@NoArgsConstructor
public class RabbitMQProducer implements Producer {
    @Autowired
    private AmqpTemplate template;
    @Autowired
    private Queue queue;

    private LocalDateTime timestamp;
    @Scheduled(fixedDelay = 1000)
    public void publish() {
        timestamp = LocalDateTime.now();
        MessageDTO messageDTO = new MessageDTO("Hello world", timestamp);
        this.template.convertAndSend(this.queue.getName(), messageDTO);
    }
}
