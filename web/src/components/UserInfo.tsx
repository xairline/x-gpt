import React, {useEffect, useState} from "react";
import {Avatar, Button, Col, Row, Typography} from "antd";
import {useLogto, UserInfoResponse} from '@logto/react';
import {LogoutOutlined} from "@ant-design/icons";
import useBreakpoint from "antd/es/grid/hooks/useBreakpoint";

const {Text,} = Typography;

interface Props {
}

const UserInfo: React.FC<Props> = (props) => {
    const {isAuthenticated, signIn, signOut, fetchUserInfo, getAccessToken} = useLogto();
    const screens = useBreakpoint();
    const [userMetadata, setUserMetadata] = useState<UserInfoResponse>();
    useEffect(() => {
        (async () => {
            if (isAuthenticated) {
                const userInfo = await fetchUserInfo();
                setUserMetadata(userInfo);
            }
        })();

    }, [isAuthenticated, fetchUserInfo]);

    return ((
        <Row>
            {userMetadata && isAuthenticated ? (
                <Row style={{
                    background: "white",
                    // paddingLeft: "12px",
                    // paddingRight: "16px",
                    borderRadius: "24px",
                }}>
                    <Col span={3} offset={1}>
                        <Avatar src={"https://xsgames.co/randomusers/avatar.php?g=pixel&key=1"}
                                size={36}
                                style={{marginRight: "12px"}}/>
                    </Col>
                    <Col span={14} offset={1}>
                        <div style={{
                            overflow: "hidden",
                            whiteSpace: "nowrap",
                            textOverflow: "ellipsis",
                            width: screens.lg ? "300px" : "200px",
                        }}>
                            <Text style={{
                                overflow: "hidden",
                                whiteSpace: "nowrap",
                                textOverflow: "ellipsis",
                                width: "100%",
                            }}>
                                {userMetadata.email}
                            </Text>
                        </div>
                    </Col>
                    <Col span={4}>
                        <Button onClick={() => signOut(`${window.location.origin}/`)}><LogoutOutlined/>
                        </Button>
                    </Col>


                </Row>
            ) : (
                <>
                    <Col span={8}></Col>
                    <Col span={8}></Col>
                    <Col span={8}>
                        <Button
                            onClick={() => signIn(`${window.location.origin}/callback`)}>Sign In
                        </Button>
                    </Col>
                </>
            )
            }
        </Row>
    ))
        ;
};

export default UserInfo;
