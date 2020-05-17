import fetch from 'isomorphic-unfetch'
import {GetStaticProps, NextPage} from 'next'
import Head from 'next/head'

import Layout from '../../components/layout'
import { useRouter } from 'next/dist/client/router'
import { AnswerType } from '../../types/questions'
import SubmitAnswer from '../../components/submit'


type InitialProps = {
  id: string | string[];
  result: AnswerType;
};

const Result: NextPage<InitialProps> = props => {
  const router = useRouter()
  console.log(props)
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
        <div>{props.result.result}</div>
        {(props.result.result !== "AC")?<pre>{props.result.error}</pre>:<div></div>}

    </Layout>
  )
}


Result.getInitialProps = async function(context){
  const { id } = context.query
  const res = await fetch(process.env.API_ENDPOINT + `/api/answers/${id}`)
  const result = await res.json() as AnswerType
  const props :InitialProps = {
    id: id,
    result: result
  }
  return props
}

export default Result