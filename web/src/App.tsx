import React, {useEffect, useState} from 'react';
import 'antd/dist/reset.css';
// import "@ant-design/plots/dist/index.css";
import './App.css';
import {Col, ConfigProvider, Image, Layout, Row, Typography} from 'antd';
import UserInfo from "./components/UserInfo";
import {Route, Routes} from "react-router-dom";
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
                            <Image src={"/logo512.png"}
                                   style={{maxHeight: "8vh", objectFit: "contain", margin: "12px 24px 12px"}}>
                            </Image>

                        </Row>
                    </Col>
                    {
                        !screens.sm ? <></> :
                            <Col flex={"auto"}>
                                <Title level={1} style={{
                                    color: "white", display: "flex",
                                    justifyContent: "center",
                                    alignItems: "center",
                                    height: "100%",
                                }}>
                                    X Airline
                                </Title>
                            </Col>
                    }
                    <Col flex={"auto"}>
                        <div style={{
                            display: "flex",
                            justifyContent: !screens.sm ? "center" : "flex-end",
                            alignItems: "center",
                            height: "100%",
                            marginRight: !screens.sm ? 0 : "24px",
                        }}>
                            <UserInfo/>
                        </div>
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
