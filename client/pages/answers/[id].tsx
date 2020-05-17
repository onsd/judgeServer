import fetch from 'isomorphic-unfetch'
import {GetStaticProps, NextPage} from 'next'
import Head from 'next/head'

import Layout from '../../components/layout'
import { useRouter } from 'next/dist/client/router'
import { AnswerType } from '../../types/questions'
import SubmitAnswer from '../../components/submit'


type InitialProps = {
  id: number;
  result: AnswerType;
};

const Result: NextPage<InitialProps> = props => {
  const router = useRouter()
  if(router.isFallback){
    return <Layout home={false}><div>Loading...</div></Layout>
  }
  return (
    <Layout home={false}>
      <Head>
        {/* {props.result.ID} */}
      </Head>
        <div>問題番号:{props.result.question_id}</div>  
        <h3>あなたの提出</h3>
        <pre>{props.result.answer}</pre>
        <h3>ステータス</h3>
        <div>{props.result.status}</div>
        {(props.result.status == "SUBMIT")?<a href="javascript:location.reload();">更新する</a>:<div></div>}
        <h3>結果</h3>
        {props.result.result}
        {(props.result.result == "WA")?<div>{props.result.detail}</div>:<div></div>}

    </Layout>
  )
}


// 最初に実行される。事前ビルドするパスを配列でreturnする。
export async function getStaticPaths() {
    // zeitが管理するレポジトリを(APIのデフォルトである)30件取得する
    const res = await fetch(process.env.API_ENDPOINT + '/api/answers')
    const repos = await res.json() as AnswerType[]
    // // レポジトリの名前をパスとする
    const paths = repos.map(repo => `/answers/${repo.ID}`)
    // const paths = [`/answers/1`,`/answers/2`]
    // 事前ビルドしたいパスをpathsとして渡す fallbackについては後述
    return { paths, fallback: true }
}

export const getStaticProps: GetStaticProps = async context => {
  const id = context.params.id
  const res = await fetch(process.env.API_ENDPOINT + `/api/answers/${id}`)
  const result = await res.json() as AnswerType
  return {props:{id, result}}
}

export default Result