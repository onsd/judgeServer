import Layout from "../components/layout";
import Head from 'next/head'
import utilStyles from '../styles/utils.module.css'
import * as React from 'react'
import {useState} from 'react'
import {  Case, QuestionType } from "../types/questions";
import fetch from 'isomorphic-unfetch'
import { useRouter } from "next/router";

// export default function createQuestion() {
//     return(
//         <Layout home={false}>
//             <Head>
//                 <title>問題を作る</title>
//             </Head>
//             <section className={utilStyles.headingMd}>
//                 <p>問題を投稿してください。</p>

//             </section>
//             <section className={`${utilStyles.headingMd} ${utilStyles.padding1px}`}>
//                 <h2 className={utilStyles.headingLg}>問題</h2>
//                 <ul className={utilStyles.list}>
                
//                 </ul>
//             </section>
//         </Layout>
//     )
// }

// title:      string;
//     body:       string;
//     validation: string;
//     input:      string;
//     output:     string;
//     testcase:   Case[];
//     samplecase: Case[];


const escapeHTML = (str:string) => {
    str = str.replace(/&/g, '&amp;');
    str = str.replace(/</g, '&lt;');
    str = str.replace(/>/g, '&gt;');
    str = str.replace(/"/g, '&quot;');
    str = str.replace(/'/g, '&#39;');
    return str;
}

const createQuestion: React.FC = () => {
    const router = useRouter()
    const [title, setTitle] = useState<string>("")
    const [body, setBody] = useState<string>("")
    const [validation, setValidation] = useState<string>("")
    const [input, setInput] = useState<string>("")
    const [output, setOutput] = useState<string>("")
    const [testcase, setTestcase] = useState<Case[]>([{Input:"", Output:""}])
    const [samplecase, setSamplecase] = useState<Case[]>([{Input:"", Output:""}])

    const updateArray = (cases: Case[], index:number, type:string, value:string, setState: React.Dispatch<React.SetStateAction<Case[]>>) => {
        const copy = cases.slice()
        copy[index][type] = value
        setState(copy)
    }

    const submitNewQuestion = (e) => {
        console.log(title,body,input,output,testcase,samplecase)
        let flag = false
        if(title == "" || body == "" || input == "" || output == "" ) {
            alert("抜けがあります")
            flag = true
        }else{
            const sanitizedTestcase = testcase.map(i => {
                if(i.Input == "" || i.Output == ""){
                    alert("抜けがあります")
                    flag = true
                }
                return {
                    Input: escapeHTML(i.Input),
                    Output: escapeHTML(i.Output)
                }
            })
            const sanitizedSamplecase = samplecase.map(i => {
                if(i.Input == "" || i.Output == ""){
                    alert("抜けがあります")
                    flag = true
                }
                return {
                    Input: escapeHTML(i.Input),
                    Output: escapeHTML(i.Output)
                }
            })
            if(flag){
                return
            }
            const question: QuestionType = {
                title: escapeHTML(title),
                body: escapeHTML(body),
                validation: validation,
                input: escapeHTML(input),
                output: escapeHTML(output),
                samplecase: sanitizedSamplecase,
                testcase: sanitizedTestcase,
            }
            fetch(process.env.API_ENDPOINT + "/api/questions",
            {
            method: "POST",
            headers: {
            "Content-Type": "application/json; charset=utf-8",
            'Accept': 'application/json'
            },
            body: JSON.stringify(question)
        }).then(res => {
            return res.json()
        }).then(json => {
            console.log(json)
            router.push("/questions/" + json.ID)
        })
        }

    }
    return(
        <Layout home={false}>
            <Head>
                <title>問題を作る</title>
            </Head>
            <section className={utilStyles.headingMd}>
            <p className={utilStyles.lightText}>プレーンテキストのみ投稿可能です。</p>
            {/* <p className={utilStyles.lightText}>HT</p> */}

            </section>
            <section className={`${utilStyles.headingMd} ${utilStyles.padding1px}`}>
                <h2 className={utilStyles.headingLg}>タイトル</h2>
                <textarea name="body" id="" cols={100} rows={5} value={title} onChange={e=>setTitle(e.target.value)}/>
                <h2 className={utilStyles.headingLg}>問題 </h2>
                <textarea name="body" id="" cols={100} rows={15} value={body} onChange={e=>setBody(e.target.value)}></textarea>
                <h2 className={utilStyles.headingLg}>制約</h2>
                <textarea name="input" id="" cols={100} rows={4} value={validation} onChange={e=>setValidation(e.target.value)}/>
                <h2 className={utilStyles.headingLg}>入力例</h2>
                <textarea name="input" id="" cols={100} rows={4} value={input} onChange={e=>setInput(e.target.value)}/>
                <h2 className={utilStyles.headingLg}>出力例</h2>
                <textarea name="title" id="" cols={100} rows={4} value={output} onChange={e=>setOutput(e.target.value)}/>
                {
                samplecase.map((value,index) => 
                    <div key={index+1}>
                        <h3 key={"samplecase-title-"+index.toString()}>サンプルケース {index+1}</h3>
                        <h3 key={"samplecase-input-name-"+index.toString()} className={utilStyles.headingLg}>入力</h3>
                        <textarea name="input" key={"samplecase-input-"+index.toString()} cols={100} rows={5} 
                            value={samplecase[index].Input} 
                            onChange={e=>updateArray(samplecase, index, "Input", e.target.value, setSamplecase)}/>
                        <h3 key={"samplecase-output-name-"+index.toString()} className={utilStyles.headingLg}>出力</h3>  
                        <textarea name="title" key={"samplecase-output-"+index.toString()} cols={100} rows={5} 
                            value={samplecase[index].Output} 
                            onChange={e=>updateArray(samplecase, index, "Output", e.target.value, setSamplecase)}/>
                    </div>      
                )
                }
                <a onClick={e=>setSamplecase([...samplecase, {Input:"", Output:""}])}>サンプルケースを増やす</a><br/>
                <a onClick={e=>setSamplecase(samplecase.slice(0,samplecase.length-1))}>サンプルケースを減らす</a>


                {
                testcase.map((value,index) => 
                    <div key={index+1}>
                        <h3 key={"testcase-title-"+index.toString()}>ジャッジ用テストケース {index+1}</h3>
                        <h3 key={"testcase-input-name-"+index.toString()} className={utilStyles.headingLg}>入力</h3>
                        <textarea name="input" key={"testcase-input-"+index.toString()} cols={100} rows={4} 
                            value={testcase[index].Input} 
                            onChange={e=>updateArray(testcase, index, "Input", e.target.value, setTestcase)}/>
                        <h3 key={"testcase-output-name-"+index.toString()} className={utilStyles.headingLg}>出力</h3>  
                        <textarea name="title" key={"testcase-output-"+index.toString()} cols={100} rows={4} 
                            value={testcase[index].Output} 
                            onChange={e=>updateArray(testcase, index, "Output", e.target.value, setTestcase)}/>
                    </div>      
                )
                }
                <a onClick={e=>setTestcase([...testcase, {Input:"", Output:""}])}>テストケースを増やす</a><br/>
                <a onClick={e=>setTestcase(testcase.slice(0,testcase.length-1))}>テストケースを減らす</a>
<br/>



                <button onClick={submitNewQuestion} >送信</button>
            </section>
        </Layout>
    )
}

export default createQuestion