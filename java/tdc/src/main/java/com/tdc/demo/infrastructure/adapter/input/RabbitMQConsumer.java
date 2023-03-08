package com.tdc.demo.infrastructure.adapter.input;

import org.springframework.amqp.rabbit.annotation.RabbitHandler;
import org.springframework.amqp.rabbit.annotation.RabbitListener;
import org.springframework.stereotype.Component;

import com.tdc.demo.domain.boundary.input.Consumer;
import com.tdc.demo.domain.dto.MessageDTO;

@Component
@RabbitListener(queues = "${rabbitmq.queue}", id = "listener")
public class RabbitMQConsumer implements Consumer{
    @RabbitHandler
    public void subscribe(MessageDTO messageDTO) {
        System.out.println("Received message. Content: " +
                messageDTO.getContent() + " -- timestamp: " + messageDTO.getTimestamp()
        );
    }
}
