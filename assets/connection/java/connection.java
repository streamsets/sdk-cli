package com.streamsets.pipeline.stage.common.@project@;

import com.streamsets.pipeline.api.ConfigDef;
import com.streamsets.pipeline.api.ConfigGroups;
import com.streamsets.pipeline.api.ConnectionDef;
import com.streamsets.pipeline.api.ConnectionEngine;
import com.streamsets.pipeline.api.Dependency;
import com.streamsets.pipeline.api.InterfaceAudience;
import com.streamsets.pipeline.api.InterfaceStability;
import com.streamsets.pipeline.api.ValueChooserModel;

import static com.streamsets.pipeline.stage.common.@project@.@name@ConnectionGroupConstants.EXAMPLE;
import static com.streamsets.pipeline.stage.common.@project@.@name@ConnectionGroupConstants.ADVANCED;

@InterfaceAudience.LimitedPrivate
@InterfaceStability.Unstable
@ConnectionDef(
    label = "@name@",
    type = @name@Connection.TYPE,
    description = "Execute Example",
    version = 1,
    upgraderDef = "upgrader/@name@ConnectionUpgrader.yaml",
    supportedEngines = {ConnectionEngine.COLLECTOR }
)
@ConfigGroups(@name@ConnectionGroups.class)

public class @name@Connection {
  public static final String TYPE = "STREAMSETS_@name@";

  @ConfigDef(
      type = ConfigDef.Type.STRING,
      label = "Example Name",
      description = "Example of connection parameter",
      required = true,
      defaultValue = "",
      group = TAP,
      displayPosition = 10,
      displayMode = ConfigDef.DisplayMode.BASIC
  )
  public String exampleName;


}
