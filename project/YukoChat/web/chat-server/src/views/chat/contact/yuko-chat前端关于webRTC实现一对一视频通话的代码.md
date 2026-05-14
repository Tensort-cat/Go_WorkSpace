```javascript
const showAVContainerModal = () => {
      data.isAVContainerModalVisible = true;
    };

    const closeAVContainerModal = () => {
      if (data.localVideo || data.remoteVideo) {
        ElMessage.warning("请先结束通话");
        return;
      }
      data.isAVContainerModalVisible = false;
    };

    const createRtcPeerConnection = () => {
      if (data.rtcPeerConn) {
        console.log("peer connection has already been created.");
        return;
      }
      data.rtcPeerConn = new RTCPeerConnection(data.ICE_CFG);
      data.rtcPeerConn.onicecandidate = (event) => {
        if (event.candidate) {
          var proxyCandidateMessage = {
            messageId: "PROXY",
            type: "candidate",
            messageData: {
              candidate: event.candidate,
            },
          };
          const rtcMessageRequest = {
            session_id: data.sessionId,
            type: 3,
            content: "",
            url: "",
            send_id: data.userInfo.uuid,
            send_name: data.userInfo.nickname,
            send_avatar: data.userInfo.avatar,
            receive_id: data.contactInfo.contact_id,
            file_size: "",
            file_name: "",
            file_type: "",
            av_data: JSON.stringify(proxyCandidateMessage),
          };
          console.log(rtcMessageRequest);
          store.state.socket.send(JSON.stringify(rtcMessageRequest));
        }
      };
      data.rtcPeerConn.oniceconnectionstatechange = (event) => {
        console.log(
          "oniceconnectionstatechange",
          data.rtcPeerConn.iceConnectionState
        );
      };
      // 对端传来媒体轨道
      data.rtcPeerConn.ontrack = (event) => {
        if (data.remoteStream === null) {
          data.remoteStream = new MediaStream();
          data.remoteVideo = document.querySelector("video.remote-video");
          data.remoteVideo.srcObject = data.remoteStream;
          data.remoteVideo.style.display = "inline-block";
        }
        data.remoteStream.addTrack(event.track);
      };
    };

    const closeRtcPeerConnection = () => {
      if (data.rtcPeerConn) {
        data.rtcPeerConn.close();
        data.rtcPeerConn = null;
      }
    };

    const getLocalMediaStream = async () => {
      if (!navigator.mediaDevices || !navigator.mediaDevices.getUserMedia) {
        console.log("getUserMedia is not supported!");
        return null;
      }

      if (data.localStream) {
        console.log("localStream already exist.");
        return data.localStream;
      }

      try {
        const stream = await navigator.mediaDevices.getUserMedia({
          video: true,
          audio: true,
        });
        return stream;
      } catch (err) {
        console.error("getUserMedia failed:", err);
        return null;
      }
    };

    const closeLocalMediaStream = () => {
      if (data.localStream != null) {
        data.localStream.getTracks().forEach((track) => {
          track.stop();
        });
        data.localStream = null;
      }
    };

    const attachMediaStreamToLocalVideo = () => {
      data.localVideo = document.querySelector("video.local-video");
      data.localVideo.srcObject = data.localStream;
      data.localVideo.muted = true;
      data.localVideo.style.display = "inline-block";
    };

    const attachMediaStreamToPeerConnection = () => {
      if (!data.localStream) {
        console.error("localStream is null!");
        return;
      }

      data.localStream.getTracks().forEach((track) => {
        data.rtcPeerConn.addTrack(track, data.localStream);
      });
    };

    const createOffer = () => {
      var offerOpts = {
        offerToReceiveAudio: true,
        offerToReceiveVideo: true,
      };
      data.rtcPeerConn
        .createOffer(offerOpts)
        .then((desc) => {
          data.rtcPeerConn.setLocalDescription(desc);
          var proxySdpMessage = {
            messageId: "PROXY",
            type: "sdp",
            messageData: {
              sdp: desc,
            },
          };
          console.log(desc);
          const rtcMessageRequest = {
            session_id: data.sessionId,
            type: 3,
            content: "",
            url: "",
            send_id: data.userInfo.uuid,
            send_name: data.userInfo.nickname,
            send_avatar: data.userInfo.avatar,
            receive_id: data.contactInfo.contact_id,
            file_size: "",
            file_name: "",
            file_type: "",
            av_data: JSON.stringify(proxySdpMessage),
          };
          store.state.socket.send(JSON.stringify(rtcMessageRequest));
        })
        .catch((err) => {
          console.log(
            `createOffer failed, error name: ${err.name}, error message: ${err.message}`
          );
        });
    };

    const createAnswer = () => {
      data.rtcPeerConn
        .createAnswer()
        .then((desc) => {
          data.rtcPeerConn.setLocalDescription(desc);
          var proxySdpMessage = {
            messageId: "PROXY",
            type: "sdp",
            messageData: {
              sdp: desc,
            },
          };
          console.log(desc);
          const rtcMessageRequest = {
            session_id: data.sessionId,
            type: 3,
            content: "",
            url: "",
            send_id: data.userInfo.uuid,
            send_name: data.userInfo.nickname,
            send_avatar: data.userInfo.avatar,
            receive_id: data.contactInfo.contact_id,
            file_size: "",
            file_name: "",
            file_type: "",
            av_data: JSON.stringify(proxySdpMessage),
          };
          store.state.socket.send(JSON.stringify(rtcMessageRequest));
        })
        .catch((err) => {
          console.log(
            `createAnswer failed, error name: ${err.name}, error message: ${err.message}`
          );
        });
    };

    const startCall = async (isInitiator) => {
      console.log(data.localVideo);
      console.log(data.localStream);
      if (data.localVideo) {
        ElMessage.warning("已经在通话中，请勿重复发起");
        return;
      }
      if (isInitiator && !data.ableToStartCall) {
        ElMessage.warning(
          "对方已经发起通话，请先接收通话或拒绝通话，才能发起下一次通话"
        );
        return;
      }
      if (!isInitiator && !data.ableToReceiveOrRejectCall) {
        ElMessage.warning("对方没有发起通话或已在通话中，无法接收通话");
        return;
      }
      createRtcPeerConnection();
      data.localStream = await getLocalMediaStream();
      if (!data.localStream) {
        ElMessage.error("无法获取摄像头/麦克风权限");
        return;
      }
      attachMediaStreamToLocalVideo();
      attachMediaStreamToPeerConnection();
      if (isInitiator) {
        var startCallMessage = {
          messageId: "PROXY",
          type: "start_call",
        };
        const rtcMessageRequest = {
          session_id: data.sessionId,
          type: 3,
          content: "",
          url: "",
          send_id: data.userInfo.uuid,
          send_name: data.userInfo.nickname,
          send_avatar: data.userInfo.avatar,
          receive_id: data.contactInfo.contact_id,
          file_size: "",
          file_name: "",
          file_type: "",
          av_data: JSON.stringify(startCallMessage),
        };
        store.state.socket.send(JSON.stringify(rtcMessageRequest));
      } else {
        var receiveCallMessage = {
          messageId: "PROXY",
          type: "receive_call",
        };
        const rtcMessageRequest = {
          session_id: data.sessionId,
          type: 3,
          content: "",
          url: "",
          send_id: data.userInfo.uuid,
          send_name: data.userInfo.nickname,
          send_avatar: data.userInfo.avatar,
          receive_id: data.contactInfo.contact_id,
          file_size: "",
          file_name: "",
          file_type: "",
          av_data: JSON.stringify(receiveCallMessage),
        };
        store.state.socket.send(JSON.stringify(rtcMessageRequest));
        data.ableToReceiveOrRejectCall = false;
      }
    };

    const sendEndCall = () => {
      if (data.localVideo == null && data.remoteVideo == null) {
        ElMessage.warning("尚未开始通话，无法挂断");
        return;
      }
      if (data.localVideo) data.localVideo.style.display = "none";
      if (data.remoteVideo) data.remoteVideo.style.display = "none";
      closeLocalMediaStream();
      closeRtcPeerConnection();
      data.remoteStream = null;
      data.localStream = null;
      data.localVideo = null;
      data.remoteVideo = null;
      data.ableToReceiveOrRejectCall = false;
      data.ableToStartCall = true;
      var proxyPeerLeaveMessage = {
        messageId: "PEER_LEAVE",
      };
      const rtcMessageRequest = {
        session_id: data.sessionId,
        type: 3,
        content: "",
        url: "",
        send_id: data.userInfo.uuid,
        send_name: data.userInfo.nickname,
        send_avatar: data.userInfo.avatar,
        receive_id: data.contactInfo.contact_id,
        file_size: "",
        file_name: "",
        file_type: "",
        av_data: JSON.stringify(proxyPeerLeaveMessage),
      };
      store.state.socket.send(JSON.stringify(rtcMessageRequest));
    };

    const endCall = () => {
      if (data.localVideo) data.localVideo.style.display = "none";
      if (data.remoteVideo) data.remoteVideo.style.display = "none";
      closeLocalMediaStream();
      closeRtcPeerConnection();
      data.remoteStream = null;
      data.localStream = null;
      data.localVideo = null;
      data.remoteVideo = null;
      data.ableToReceiveOrRejectCall = false;
      data.ableToStartCall = true;
      ElMessage.warning("对方拒绝通话");
    };

    const receiveEndCall = () => {
      if (data.localVideo) data.localVideo.style.display = "none";
      if (data.remoteVideo) data.remoteVideo.style.display = "none";
      closeLocalMediaStream();
      closeRtcPeerConnection();
      data.remoteStream = null;
      data.localStream = null;
      data.localVideo = null;
      data.remoteVideo = null;
      data.ableToReceiveOrRejectCall = false;
      data.ableToStartCall = true;
      ElMessage.warning("对方已挂断");
    };

    const handleOfferSdp = (val) => {
      data.rtcPeerConn
        .setRemoteDescription(new RTCSessionDescription(val))
        .then(() => {
          createAnswer();
        })
        .catch((err) => {
          console.log("rtcPeerConn setRemoteDescription failed", err);
        });
    };

    const handleAnswerSdp = (val) => {
      data.rtcPeerConn
        .setRemoteDescription(new RTCSessionDescription(val))
        .catch((err) => {
          console.log("rtcPeerConn setRemoteDescription failed", err);
        });
    };

    const handleCandidate = (val) => {
      data.rtcPeerConn.addIceCandidate(new RTCIceCandidate(val));
    };

    const rejectCall = () => {
      if (!data.ableToReceiveOrRejectCall) {
        ElMessage.warning("对方没有发起通话或已在通话中，无法拒绝通话");
        return;
      }
      var rejectCallMessage = {
        messageId: "PROXY",
        type: "reject_call",
      };
      const rtcMessageRequest = {
        session_id: data.sessionId,
        type: 3,
        content: "",
        url: "",
        send_id: data.userInfo.uuid,
        send_name: data.userInfo.nickname,
        send_avatar: data.userInfo.avatar,
        receive_id: data.contactInfo.contact_id,
        file_size: "",
        file_name: "",
        file_type: "",
        av_data: JSON.stringify(rejectCallMessage),
      };
      store.state.socket.send(JSON.stringify(rtcMessageRequest));
      data.ableToReceiveOrRejectCall = false;
    };
```

