package com.streamsets.pipeline.stage.common.@project@;

import com.streamsets.pipeline.api.Label;

public enum @name@ConnectionGroups implements Label {
  EXAMPLE("Example"),
  ADVANCED("Advanced");

  private final String label;

  @name@ConnectionGroups(String label) {
    this.label = label;
  }

  public String getLabel() {
    return this.label;
  }
}
