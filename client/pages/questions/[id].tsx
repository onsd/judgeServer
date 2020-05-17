import fetch from 'isomorphic-unfetch'
import {GetStaticProps, NextPage} from 'next'
import Head from 'next/head'
import utilStyles from '../../styles/utils.module.css'

import Layout from '../../components/layout'
import { useRouter } from 'next/dist/client/router'
import { QuestionType } from '../../types/questions'
import SubmitAnswer from '../../components/submit'


type InitialProps = {
  id: number;
  question: QuestionType;
};

const Question: NextPage<InitialProps> = props => {
  const router = useRouter()



  if(router.isFallback){
    return <Layout home={false}><div>Loading...</div></Layout>
  }
  return (
    <Layout home={false}>
      <Head>
        {props.question.ID}
      </Head>
        <div>問題番号:{props.question.ID}</div>  
        <h3>問題文</h3>
        <div>{props.question.body}</div>
        <h3>制約</h3>
        <div>{props.question.validation}</div>
        <h3>入力</h3>
        <div>{props.question.input}</div>
        <h3>出力</h3>
        <div>{props.question.output}</div>
        {
          props.question.samplecase.map((value,index) => 
              <div key={index+1}>
                  <h3 key={"samplecase-title-"+index.toString()}>サンプルケース {index}</h3>
                  <h3 key={"samplecase-input-name-"+index.toString()} >入力</h3>
                  {value.Input}
                  <h3 key={"samplecase-output-name-"+index.toString()} >出力</h3>  
                  {value.Output}
              </div>      
          )
          }
        <SubmitAnswer questionID={props.question.ID}/>
    </Layout>
  )
}


// 最初に実行される。事前ビルドするパスを配列でreturnする。
export async function getStaticPaths() {
    // zeitが管理するレポジトリを(APIのデフォルトである)30件取得する
    const res = await fetch(process.env.API_ENDPOINT + '/api/questions')
    const repos = await res.json() as QuestionType[]
    // レポジトリの名前をパスとする
    const paths = repos.map(repo => `/questions/${repo.ID}`)
    // 事前ビルドしたいパスをpathsとして渡す fallbackについては後述
    return { paths, fallback: true }
}

export const getStaticProps: GetStaticProps = async context => {
  const id = context.params.id
  const res = await fetch(process.env.API_ENDPOINT + `/api/questions/${id}`)
  const question = await res.json() as QuestionType
  return {props:{id, question}}
}

export default Question