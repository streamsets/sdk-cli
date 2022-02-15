package com.streamsets.pipeline.stage.origin.@project@;

import com.streamsets.pipeline.api.ConfigDefBean;
import com.streamsets.pipeline.api.ConfigGroups;
import com.streamsets.pipeline.api.ExecutionMode;
import com.streamsets.pipeline.api.GenerateResourceBundle;
import com.streamsets.pipeline.api.Source;
import com.streamsets.pipeline.api.StageDef;
import com.streamsets.pipeline.api.base.configurablestage.DSource;
import com.streamsets.pipeline.lib.event.NoMoreDataEvent;
import com.streamsets.pipeline.lib.@project@.config.@name@CoreConfigBean;
import com.streamsets.pipeline.lib.@project@.config.@name@Groups;

@StageDef(
    version = 1,
    label = "Example Origin",
    description = "Execute Example Origin",
    icon = "@name@.png",
    execution = ExecutionMode.STANDALONE,
    recordsByRef = true,
    producesEvents = true,
    eventDefs = {NoMoreDataEvent.class},
    onlineHelpRefUrl = "",
    upgraderDef = "upgrader/@name@DSource.yaml",
    resetOffset = true
)
@ConfigGroups(@name@Groups.class)
@GenerateResourceBundle

public class @name@DSource extends DSource {
  @ConfigDefBean()
  public @name@CoreConfigBean configBean;

  @Override
  protected Source createSource() {
    return new @name@Source(configBean);
  }
}
