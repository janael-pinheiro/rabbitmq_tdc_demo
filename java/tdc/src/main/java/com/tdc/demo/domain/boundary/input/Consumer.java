package com.tdc.demo.domain.boundary.input;

import com.tdc.demo.domain.dto.MessageDTO;

public interface Consumer {
    void subscribe(MessageDTO messageDTO);
}
