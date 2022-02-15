package com.streamsets.pipeline.lib.@project@.config;

import com.streamsets.pipeline.api.GenerateResourceBundle;
import com.streamsets.pipeline.api.Label;

@GenerateResourceBundle
public enum @name@Groups implements Label {
  EXAMPLE("Example"),
  ADVANCED("Advanced");

  private final String label;

  @name@Groups(String label) {
    this.label = label;
  }

  public String getLabel() {
    return this.label;
  }
}
