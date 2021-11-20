// @ts-check
import { AppBar, Button, makeStyles, Tab, Tabs, Typography } from "@material-ui/core";
import React from "react";
import { pSBCr } from "../../utils/lightenOrDarkenColor";
import { getMeshProperties } from "../../utils/nameMapper";
import PatternServiceFormCore from "./PatternServiceFormCore";

const useStyles = makeStyles(() => ({
  appBar : {
    // @ts-ignore
    boxShadow : ({ color }) =>  `0px 2px 4px -1px ${pSBCr(color, -30)}`,
    // @ts-ignore
    background : ({ color }) => `linear-gradient(115deg, ${pSBCr( color, -30)} 0%, ${color} 100%)`,
    position : "sticky",
  },
  tabPanel : {
    marginTop : "1.1rem",
    padding : "0px 2px"
  }
}));

function RJSFButton({ handler, text, ...restParams }) {
  return (
    <Button variant="contained" color="primary" style={{ margin : "0px 0.5rem 32px 0px" }} onClick={handler} {...restParams}>
      {text}
    </Button>
  );
}

function RJSFFormChildComponent({ onSubmit, onDelete }){
  return (
    <>
      <RJSFButton handler={onSubmit} text="Submit" />
      <RJSFButton handler={onDelete} text="Delete" />
    </>
  )
}

/**
 * PatternServiceForm renders a form from the workloads schema and
 * traits schema
 * @param {{
 *  schemaSet: { workload: any, traits: any[], type: string };
 *  onSubmit: Function;
 *  onDelete: Function;
 *  namespace: string;
 *  onChange?: Function
 *  onSettingsChange?: Function;
 *  onTraitsChange?: Function;
 *  formData?: Record<String, unknown>
 *  reference?: Record<any, any>;
 * }} props
 * @returns
 */
function PatternServiceForm({ formData, schemaSet, onSubmit, onDelete, reference, namespace, onSettingsChange, onTraitsChange }) {
  const [tab, setTab] = React.useState(0);
  const classes = useStyles({ color : getMeshProperties(getMeshName(schemaSet))?.color });

  const handleTabChange = (_, newValue) => {
    setTab(newValue);
  };
  const renderTraits = () => !!schemaSet.traits?.length;

  return (
    <PatternServiceFormCore
      formData={formData}
      schemaSet={schemaSet}
      onSubmit={onSubmit}
      onDelete={onDelete}
      reference={reference}
      namespace={namespace}
      onSettingsChange={onSettingsChange}
      onTraitsChange={onTraitsChange}
    >
      {(SettingsForm, TraitsForm) => {
        return (
          <div>
            <AppBar className={classes.appBar}>
              <Tabs value={tab} onChange={handleTabChange} aria-label="Pattern Service" >
                <Tab label="Settings" {...a11yProps(0)} />
                {
                  renderTraits()
                    ? <Tab label="Traits" {...a11yProps(1)} />
                    : null
                }
              </Tabs>
            </AppBar>
            <TabPanel value={tab} index={0} className={classes.tabPanel}>
              <SettingsForm RJSFFormChildComponent={RJSFFormChildComponent} />
            </TabPanel>
            <TabPanel value={tab} index={0} className={classes.tabPanel}>
              <TraitsForm />
            </TabPanel>
          </div>
        )
      }}
    </PatternServiceFormCore>
  )
}

function TabPanel(props) {
  const {
    children, value, index, ...other
  } = props;

  return (
    <div
      role="tabpanel"
      hidden={value !== index}
      id={`simple-tabpanel-${index}`}
      aria-labelledby={`simple-tab-${index}`}
      {...other}
    >
      {value === index && (
        <Typography>{children}</Typography>
      )}
    </div>
  );
}

function a11yProps(index) {
  return {
    id : `simple-tab-${index}`,
    "aria-controls" : `simple-tabpanel-${index}`,
  };
}

/**
 * @param {{ workload: { [x: string]: string; }; }} schema
 * @returns {String} name
 */
function getMeshName(schema) {
  return schema?.workload?.["service-mesh"]?.toLowerCase() || "core";
}

export default PatternServiceForm;