package com.streamsets.pipeline.lib.@project@.config;

import com.streamsets.pipeline.api.ConfigDef;
import com.streamsets.pipeline.api.ConfigDefBean;
import com.streamsets.pipeline.api.ConfigGroups;
import com.streamsets.pipeline.api.ConnectionDef;
import com.streamsets.pipeline.api.Dependency;
import com.streamsets.pipeline.api.ValueChooserModel;
import com.streamsets.pipeline.singerio.tap.TapImpl;
import com.streamsets.pipeline.stage.common.@project@.@name@Connection;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

@ConfigGroups(@name@Groups.class)

public class @name@CoreConfigBean {


  private static final Logger LOG = LoggerFactory.getLogger(@name@CoreConfigBean.class);

  @ConfigDef(
      required = true,
      type = ConfigDef.Type.MODEL,
      connectionType = @name@Connection.TYPE,
      defaultValue = ConnectionDef.Constants.CONNECTION_SELECT_MANUAL,
      label = "Connection",
      group = @name@GroupConstants.TAP,
      displayPosition = -500
  )
  @ValueChooserModel(ConnectionDef.Constants.ConnectionChooserValues.class)
  public String connectionSelection = ConnectionDef.Constants.CONNECTION_SELECT_MANUAL;

  @ConfigDefBean(
      dependencies = {
          @Dependency(
              configName = "connectionSelection",
              triggeredByValues = ConnectionDef.Constants.CONNECTION_SELECT_MANUAL
          )
      }
  )
  public @name@Connection connection;


  @ConfigDef(
      type = ConfigDef.Type.NUMBER,
      label = "Timeout",
      description = "Max number of seconds to wait for more records",
      required = true,
      group = @name@GroupConstants.ADVANCED,
      defaultValue = "30",
      min = 5,
      displayPosition = 50,
      displayMode = ConfigDef.DisplayMode.ADVANCED
  )
  public int timeout = 30;

  @ConfigDef(
      type = ConfigDef.Type.NUMBER,
      label = "BatchSize",
      description = "Max number of records per batch",
      required = true,
      group = @name@GroupConstants.ADVANCED,
      min = 1,
      defaultValue = "1000",
      displayPosition = 60,
      displayMode = ConfigDef.DisplayMode.ADVANCED
  )
  public int batchSize = 1000;
}
