import React, { useState } from 'react';
import fetch from 'isomorphic-unfetch'
import { useRouter, Router } from "next/router";


// export default class Form1 extends React.Component {
const Form1 = (props) => {
  // constructor(props){
  //   super(props);
  //   console.log(this.props.questionID)
  //   this.state = {
  //     usstate: props.initState,
  //     answer: 'This is for a text area.'
  //   };
  //   this.onChange = this.onChange.bind(this);
  //   this.onSubmit = this.onSubmit.bind(this);
  //   this.onTextAreaChange = this.onTextAreaChange.bind(this);
  // }

  const [answer, setAnswer] = useState("")
  const [language, setLanguage] = useState("")
  const router = useRouter()
  const onAnswerChange = (e) =>{
    console.log(e.target.value);
    setAnswer(e.target.value);
  }

  const onLanguageChange = (e) =>{
    console.log(e.target.value);
    setLanguage(e.target.value);
  }

  const onSubmit = (e) =>{
    e.preventDefault();
    console.log("onSubmit");
    const data = JSON.stringify({
      language: language,
      answer : answer
    })
    console.log(data)
    fetch(process.env.API_ENDPOINT + "/api/answers/"+ props.questionID,
    {
      method: "POST",
      headers: {
       "Content-Type": "application/json; charset=utf-8",
       'Accept': 'application/json'
      },
      body: data
   }).then(res => {
      return res.json()
   }).then(json => {
     console.log(json)
     router.push("/answers/" + json.ID)
   })
    
  }

  const onTextAreaChange = (e) =>{
    this.setState({ answer: e.target.value });
  }

  // render() {
    var states = [
      { code: "", name:"選んでください"},
      { code: "pyton", name: "python" },
      { code: "cpp", name: "c++" },
      { code: "javascript", name: "JavaScript"}]
    const options = states.map(
      (n)=>(
        <option key={n.code} value={n.code}>
          {n.name}
        </option>
      ))
  //   );
    return (
      <form >
        <div>
          <select
            value={language}
            onChange={onLanguageChange}
            >
            {options}
          </select>
        </div>
        <textarea
          rows="10"
          cols="50"
          value={answer}
          onChange={onAnswerChange}/>
        <div>
          <button type="submit" onClick={onSubmit}>OK</button>
        </div>
      </form>
    );
  }

export default Form1