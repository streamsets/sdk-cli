package com.streamsets.pipeline.lib.@project@.config;

import com.streamsets.pipeline.api.ErrorCode;
import com.streamsets.pipeline.api.GenerateResourceBundle;

@GenerateResourceBundle
public enum @name@Errors implements ErrorCode {
  @name@_00("Example error message: {}"),
  @name@_01("Another error message: {}, {}")

  private final String msg;

  @name@Errors(String msg) {
    this.msg = msg;
  }

  @Override
  public String getCode() {
    return name();
  }

  @Override
  public String getMessage() {
    return msg;
  }
}
