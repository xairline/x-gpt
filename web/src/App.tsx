import React, {useEffect, useState} from 'react';
import 'antd/dist/reset.css';
// import "@ant-design/plots/dist/index.css";
import './App.css';
import {Col, ConfigProvider, Image, Layout, Modal, Row, Typography} from 'antd';
import UserInfo from "./components/UserInfo";
import {Link, Route, Routes} from "react-router-dom";
import Dashboard from "./pages/Dashboard";
import useBreakpoint from "antd/es/grid/hooks/useBreakpoint";
import {useLogto} from '@logto/react';
import Callback from "./pages/Callback";
import FlightLog from "./pages/flight-log";
import Markdown from "react-markdown";

const {Title} = Typography;
const {Content,} = Layout;

const App: React.FC = () => {
    let location = window.location.pathname;
    const screens = useBreakpoint();
    const [token, setToken] = useState("");
    const {isAuthenticated, fetchUserInfo, getAccessTokenClaims, getIdToken} = useLogto();
    const [isModalOpen, setIsModalOpen] = useState(false);

    const showModal = () => {
        setIsModalOpen(true);
    };
    const handleOk = () => {
        setIsModalOpen(false);
    }

    useEffect(() => {
        (async () => {
            if (isAuthenticated) {
                const token = await getIdToken();
                setToken(token as string);
            }
        })();
    }, [isAuthenticated, fetchUserInfo]);
    const [current, setCurrent] = useState(
        location === "/" || location === ""
            ? "dashboard"
            : location.split("/")[1]
    );

    return ((
        <ConfigProvider
            theme={{
                token: {
                    colorPrimary: '#006363'
                },
            }}
        >
            <Layout className="layout app">
                <Row style={{background: "#006363"}}>
                    <Col flex={"auto"}>
                        <Row style={{
                            display: "flex",
                            height: "100%",
                        }}>
                            <Link to={"/"}>
                                <Image src={"/logo512.png"}
                                       style={
                                           {
                                               maxHeight: "8vh",
                                               objectFit: "contain",
                                               margin: "12px 24px 12px"
                                           }
                                       }
                                       preview={false}
                                >

                                </Image>
                            </Link>

                        </Row>
                    </Col>
                    {
                        !screens.sm ? <></> :
                            <Col flex={"auto"}>
                                <Link to={"/"}>
                                    <Title level={1} style={{
                                        color: "white", display: "flex",
                                        justifyContent: "flex-start",
                                        alignItems: "center",
                                        height: "100%",
                                    }}>
                                        X Airline
                                    </Title>
                                </Link>
                            </Col>
                    }
                    <Col flex={"auto"}>
                        <Row style={{
                            display: "flex",
                            justifyContent: !screens.sm ? "center" : "flex-end",
                            alignItems: "center",
                            height: "100%",
                            marginRight: !screens.sm ? "12px" : "24px",
                            marginLeft: !screens.sm ? "12px" : "24px",
                        }}>
                            <Col span={6}>
                                <Modal
                                    title="Download"
                                    open={isModalOpen}
                                    onOk={handleOk}
                                    cancelButtonProps={{style: {display: "none"}}}
                                >
                                    < Markdown>
                                        {`
# Download and install plugin
[Download](https://github.com/xairline/xairline-v2/releases/latest/download/XWebStack.zip) and unzip the plugin to your X-Plane plugin folder.
# Access Token

\`
${token ? "CLIENT_TOKEN=" + token : "Login/Sign up to get your access token."}
\`

# Configuration
Copy above token, including **CLIENT_TOKEN=**, and past it into the **config** file under **PATH_TO_XPLANE/Resources/plugins/XWebStack** folder.

# Access GPTs
[GPTs](https://chat.openai.com/g/g-sPpIgSxow-x-airline)                      
`}
                                    </Markdown>

                                </Modal>
                                <a href={"#"} onClick={() => {
                                    setIsModalOpen(true)
                                }}>
                                    <Title
                                        level={screens.xxl ? 2 : 5}
                                        italic={true}
                                        style={{
                                            marginTop: "24px",
                                            color: "white"
                                        }}>
                                        chatGPT
                                    </Title>
                                </a>
                            </Col>
                            <Col span={10}>
                                <a href={"https://github.com/xairline/xairline-v2/releases/latest"}>
                                    <Title
                                        level={screens.xxl ? 2 : 5}
                                        italic={true}
                                        style={{
                                            marginTop: "24px",
                                            color: "white"
                                        }}>
                                        Download
                                    </Title>
                                </a>
                            </Col>
                            <Col span={8}><UserInfo/></Col>
                        </Row>
                    </Col>
                </Row>

                <Layout>
                    <Content
                        style={{
                            padding: 24,
                            minHeight: 280,
                            background: "white",
                            overflow: "hidden",
                        }}
                    >
                        <Routes>
                            <Route path={"/"} element={<Dashboard/>}/>
                            <Route path={"/dashboard"} element={<Dashboard/>}/>
                            <Route path="/callback" element={<Callback/>}/>
                            <Route
                                key={'flight-logs'}
                                path="/flight-logs/:id"
                                element={<FlightLog/>}
                            />
                        </Routes>
                    </Content>
                </Layout>
            </Layout>
        </ConfigProvider>
    ))
        ;
};

export default App;
