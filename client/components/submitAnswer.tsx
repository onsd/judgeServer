import * as React from 'react'
import { Formik } from 'formik'
import Yup from 'yup';

// import Head from 'next/head'
// import fetch from 'isomorphic-fetch'
import Select from 'react-select'


const Editor = (props) => {
  if (typeof window !== 'undefined') {
    const Ace = require('react-ace').default;
    require('ace-builds/src-noconflict/mode-'+ props.mode);
    require('ace-builds/src-noconflict/theme-'+ props.theme);

    return <Ace {...props}/>
  }

  return null;
}

type props = {
  questionID: number,
}
const languageOptions = [
  {label: "python", value:"python"},
  {label: "c++", value:"cpp"}

]

const SubmitAnswer :React.FC<props> = ({questionID}) => {
  return(
    <Formik
    initialValues={{ language: "javascript", body: "" }}
    onSubmit={values => console.log(values)}
    render={(props) => (
      <form onSubmit={props.handleSubmit}>
        <div>
          <div>言語</div>
          {/* <input
            name="language"
            value={props.values.language}
            onChange={props.handleChange}
          /> */}
          <select
            name="color"
            value={props.values.language}
            onChange={props.handleChange}
            onBlur={props.handleBlur}
            style={{ display: 'block' }}
            defaultValue="python"
          >
            {languageOptions.map(i => <option value={i.value} label={i.label} />)}
          </select>
        </div>
        <br/>
        <br/>
        <br/>
        <div>
          <div>コード</div>
          <Editor mode={props.values.language} theme="github" value={props.values.body}onChange={props.handleChange}/>
        </div>
        <button type="submit">送信</button>
      </form>
    )}
  />
  )
}

export default SubmitAnswer