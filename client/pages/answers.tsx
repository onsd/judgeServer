import Head from 'next/head'
import Layout, { siteTitle } from '../components/layout'
import utilStyles from '../styles/utils.module.css'
import Link from 'next/link'
import fetch from 'isomorphic-unfetch'
import styles from '../components/layout.module.css'



export default function Answers({ allAnswerData }) {
  return (
    <Layout home={false}>
      <Head>
        <title>{siteTitle}</title>
      </Head>
      <section className={utilStyles.headingMd}>
        <p>提出済み
        </p>
      </section>
      <section className={`${utilStyles.headingMd} ${utilStyles.padding1px}`}>
        <h2 className={utilStyles.headingLg}>Questions</h2>
        <ul className={utilStyles.list}>
        {allAnswerData.map(answers => (
            <li className={utilStyles.listItem}>
                    <Link href="/answers/[id]" as={`/answers/${answers.ID}`}>
                            <a>回答:{answers.ID}</a>
                    </Link>                
                    <ul　 key={answers.ID}>
                    <li>
                        問題の状態：{answers.status}
                    </li>
                    <li>
                        結果: {answers.result}
                    </li>
                </ul>
            </li>
        ))}

        </ul>
      </section>
      <div className={styles.backToHome}>
          <Link href="/createQuestion">
            <a>問題をつくる</a>
          </Link>
    </div>
    </Layout>
  )
}


Answers.getInitialProps = async () => {
  const res = await fetch(process.env.API_ENDPOINT + "/api/answers")
  const json = await res.json()
  return { allAnswerData: json }
}
// export async function getStaticProps() {
//   const allPostsData = getSortedPostsData()
//   return {
//     props: {
//       allPostsData
//     }
//   }
// }
