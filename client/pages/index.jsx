import Head from 'next/head'
import Layout, { siteTitle } from '../components/layout'
import utilStyles from '../styles/utils.module.css'
import Link from 'next/link'
import fetch from 'isomorphic-unfetch'
import styles from '../components/layout.module.css'



export default function Home({ allQuestionsData }) {
  return (
    <Layout home>
      <Head>
        <title>{siteTitle}</title>
      </Head>
      <section className={utilStyles.headingMd}>
        <p>腕試しに作ってみたオンラインジャッジです。</p>
        <p>
          現状Pythonのみジャッジ可能です
        </p>
      </section>
      <section className={`${utilStyles.headingMd} ${utilStyles.padding1px}`}>
        <h2 className={utilStyles.headingLg}>Questions</h2>
        <ul className={utilStyles.list}>
          {allQuestionsData.map(questions => (
            <li className={utilStyles.listItem} key={questions.ID}>
              <Link href="/questions/[id]" as={`/questions/${questions.ID}`}>
                <a>{questions.title}</a>
              </Link>
              <br />
            </li>
          ))}
        </ul>
      </section>
      <div className={styles.backToHome}>
          <Link href="/createQuestion">
            <a>問題をつくる</a>
          </Link>
          
          <Link href="/answers">
            <a>　　　　　提出を見る</a>
          </Link>
        </div>
    </Layout>
  )
}


Home.getInitialProps = async () => {
  const res = await fetch(process.env.API_ENDPOINT + '/api/questions')
  const json = await res.json()
  return { allQuestionsData: json }
}
