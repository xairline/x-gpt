import React, {useEffect, useState} from 'react';
import 'antd/dist/reset.css';
// import "@ant-design/plots/dist/index.css";
import './App.css';
import {Col, ConfigProvider, Image, Layout, Row, Typography} from 'antd';
import UserInfo from "./components/UserInfo";
import {Link, Route, Routes} from "react-router-dom";
import Dashboard from "./pages/Dashboard";
import useBreakpoint from "antd/es/grid/hooks/useBreakpoint";
import {useLogto} from '@logto/react';
import Callback from "./pages/Callback";
import FlightLog from "./pages/flight-log";

const {Title} = Typography;
const {Content,} = Layout;

const App: React.FC = () => {
    let location = window.location.pathname;
    const screens = useBreakpoint();
    const {isAuthenticated, fetchUserInfo} = useLogto();
    useEffect(() => {

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
                        }}>
                            <Col span={6}>
                                <a href={"https://docs.xairline.org"}>
                                    <Title
                                        level={screens.xxl ? 2 : 5}
                                        italic={true}
                                        style={{
                                            marginTop: "24px",
                                            color: "white"
                                        }}>
                                        Docs
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
